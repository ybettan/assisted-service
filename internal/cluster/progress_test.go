package cluster

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/openshift/assisted-service/internal/common"
	"github.com/openshift/assisted-service/internal/events"
	"github.com/openshift/assisted-service/internal/host"
	"github.com/openshift/assisted-service/internal/metrics"
	"github.com/openshift/assisted-service/internal/operators"
	"github.com/openshift/assisted-service/models"
	"github.com/openshift/assisted-service/pkg/s3wrapper"
)

var _ = Describe("Progress bar test", func() {

	var (
		ctx              = context.Background()
		db               *gorm.DB
		dbName           string
		ctrl             *gomock.Controller
		cluster          common.Cluster
		clusterApi       *Manager
		mockEvents       *events.MockHandler
		mockHostAPI      *host.MockAPI
		mockMetric       *metrics.MockAPI
		mockS3Api        *s3wrapper.MockAPI
		operatorsManager *operators.Manager
	)

	BeforeEach(func() {
		db, dbName = common.PrepareTestDB()
		ctrl = gomock.NewController(GinkgoT())
		mockEvents = events.NewMockHandler(ctrl)
		mockHostAPI = host.NewMockAPI(ctrl)
		mockMetric = metrics.NewMockAPI(ctrl)
		mockS3Api = s3wrapper.NewMockAPI(ctrl)
		operatorsManager = operators.NewManager(common.GetTestLog(), nil, operators.Options{}, nil)
		clusterApi = NewManager(getDefaultConfig(), common.GetTestLog().WithField("pkg", "cluster-monitor"), db,
			mockEvents, mockHostAPI, mockMetric, nil, nil, operatorsManager, nil, mockS3Api, nil)
	})

	AfterEach(func() {
		common.DeleteTestDB(db, dbName)
		ctrl.Finish()
	})

	expectProgressToBe := func(totalPercentage *int, preparingForInstallationStagePercentage, installingStagePercentage, finalizingStagePercentage int) {

		Expect(cluster.Progress).NotTo(BeNil())
		if totalPercentage == nil {
			Expect(cluster.Progress.TotalPercentage).To(BeNil())
		} else {
			Expect(*cluster.Progress.TotalPercentage).To(Equal(int64(*totalPercentage)))
		}
		Expect(cluster.Progress.PreparingForInstallationStagePercentage).To(Equal(int64(preparingForInstallationStagePercentage)))
		Expect(cluster.Progress.InstallingStagePercentage).To(Equal(int64(installingStagePercentage)))
		Expect(cluster.Progress.FinalizingStagePercentage).To(Equal(int64(finalizingStagePercentage)))
	}

	FIt("total_progress", func() {

		By("create the cluster", func() {

			clusterId := strfmt.UUID(uuid.New().String())
			cluster = common.Cluster{
				Cluster: models.Cluster{
					ID:                       &clusterId,
					APIVip:                   "1.2.3.5",
					IngressVip:               "1.2.3.6",
					MachineNetworkCidr:       "1.2.3.0/24",
					Status:                   swag.String(models.ClusterStatusReady),
					StatusInfo:               swag.String(StatusInfoReady),
					VipDhcpAllocation:        swag.Bool(true),
					ServiceNetworkCidr:       "1.2.4.0/24",
					ClusterNetworkCidr:       "1.3.0.0/16",
					BaseDNSDomain:            "test.invalid",
					PullSecretSet:            true,
					ClusterNetworkHostPrefix: 24,
					//FIXME:ybettan: add operators
					MonitoredOperators: []*models.MonitoredOperator{
						{
							Name:         operators.OperatorConsole.Name,
							OperatorType: models.OperatorTypeBuiltin,
							Status:       models.OperatorStatusAvailable,
						},
					},
				},
			}
			Expect(db.Create(&cluster).Error).ShouldNot(HaveOccurred())

			for i := 0; i < 5; i++ {

				id := strfmt.UUID(uuid.New().String())
				host := models.Host{
					ID:        &id,
					ClusterID: clusterId,
					Status:    swag.String(models.HostStatusKnown),
					Inventory: common.GenerateTestDefaultInventory(),
					Role:      models.HostRoleMaster,
				}

				if i > 2 {
					host.Role = models.HostRoleWorker
				}

				Expect(db.Create(&host).Error).ShouldNot(HaveOccurred())
			}

			clusterAfterRefresh, err := clusterApi.RefreshStatus(ctx, &cluster, db)
			Expect(err).To(BeNil())

			Expect(*clusterAfterRefresh.Status).To(Equal(models.ClusterStatusReady))
			expectProgressToBe(nil, 0, 0, 0)
		})

		By("preparing_for_installation_stage_percentage", func() {

			mockEvents.EXPECT().AddEvent(ctx, *cluster.ID, nil, gomock.Any(), gomock.Any(), gomock.Any())

			err := clusterApi.PrepareForInstallation(ctx, &cluster, db)
			Expect(err).NotTo(HaveOccurred())

			cluster = getClusterFromDB(*cluster.ID, db)

			Expect(*cluster.Status).To(Equal(models.ClusterStatusPreparingForInstallation))
			expectProgressToBe(nil, 0, 0, 0)
		})

		////if t.withOCMClient {
		////	mockAccountsMgmt = ocm.NewMockOCMAccountsMgmt(ctrl)
		////	ocmClient := &ocm.Client{AccountsMgmt: mockAccountsMgmt, Config: &ocm.Config{WithAMSSubscriptions: true}}
		////	clusterApi = NewManager(getDefaultConfig(), common.GetTestLog().WithField("pkg", "cluster-monitor"), db,
		////		mockEvents, mockHostAPI, mockMetric, nil, nil, operatorsManager, ocmClient, mockS3Api, nil)
		////	if !t.requiresAMSUpdate {
		////		cluster.IsAmsSubscriptionConsoleUrlSet = true
		////	}
		////}
		////cluster.MachineNetworkCidrUpdatedAt = time.Now().Add(-3 * time.Minute)
		//Expect(db.Create(&cluster).Error).ShouldNot(HaveOccurred())
		//for i := range t.hosts {
		//	t.hosts[i].ClusterID = clusterId
		//	Expect(db.Create(&t.hosts[i]).Error).ShouldNot(HaveOccurred())
		//}
		////cluster = getClusterFromDB(clusterId, db)
		////if srcState != t.dstState {
		////	mockEvents.EXPECT().AddEvent(gomock.Any(), gomock.Any(), gomock.Any(),
		////		gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
		////}
		////if t.srcState == models.ClusterStatusFinalizing && !t.requiresAMSUpdate {
		////	mockS3Api.EXPECT().DoesObjectExist(ctx, fmt.Sprintf("%s/%s", cluster.ID, constants.Kubeconfig)).Return(false, nil)
		////}
		////reportInstallationCompleteStatuses := []string{models.ClusterStatusInstalled, models.ClusterStatusError}
		////if funk.Contains(reportInstallationCompleteStatuses, t.dstState) {
		////	mockMetricsAPIInstallationFinished()
		////} else if t.dstState == models.ClusterStatusInsufficient {
		////	mockHostAPIIsRequireUserActionResetFalse()
		////}
		////if t.requiresAMSUpdate {
		////	subscriptionID := cluster.AmsSubscriptionID
		////	consoleUrl := common.GetConsoleUrl(cluster.Name, cluster.BaseDNSDomain)
		////	mockAccountsMgmt.EXPECT().UpdateSubscriptionConsoleUrl(ctx, subscriptionID, consoleUrl)
		////}
		////Expect(cluster.ValidationsInfo).To(BeEmpty())
		//clusterAfterRefresh, err := clusterApi.RefreshStatus(ctx, &cluster, db)
		//Expect(clusterAfterRefresh).ToNot(BeNil())
		////Expect(clusterAfterRefresh.ValidationsInfo).To(BeEmpty())
		////Expect(err).ToNot(HaveOccurred())
		////Expect(clusterAfterRefresh.Status).To(Equal(&t.dstState))
		////t.statusInfoChecker.check(clusterAfterRefresh.StatusInfo)
	})
})
