package main

import (
	"github.com/kr/pretty"
	rotatorModel "github.com/mattermost/rotator/model"
	"github.com/mattermost/rotator/rotator"
	"github.com/pkg/errors"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// rotateCmd represents the rotate command
var rotateCmd = &cobra.Command{
	Use:   "rotate",
	Short: "Rotates nodes",
	Long:  `Applying latest AMI that exists on launch template or launch configuration and rolls out new nodes`,
	RunE: func(cmd *cobra.Command, args []string) error {
		pretty.Println("rotate called")

		cluster, _ := cmd.Flags().GetString("cluster")
		logger := logger.WithField("cluster", cluster)
		maxScaling, _ := cmd.Flags().GetInt("max-scaling")
		rotateMasters, _ := cmd.Flags().GetBool("rotate-masters")
		rotateWorkers, _ := cmd.Flags().GetBool("rotate-workers")
		maxDrainRetries, _ := cmd.Flags().GetInt("max-drain-retries")
		evictGracePeriod, _ := cmd.Flags().GetInt("evict-grace-period")
		waitBetweenRotations, _ := cmd.Flags().GetInt("wait-between-rotations")
		waitBetweenDrains, _ := cmd.Flags().GetInt("wait-between-drains")

		if len(cluster) == 0 {
			logger.Fatal("cluster is not set.")
			return nil
		}

		clusterRotator := rotatorModel.Cluster{
			ClusterID:            cluster,
			MaxScaling:           maxScaling,
			RotateMasters:        rotateMasters,
			RotateWorkers:        rotateWorkers,
			MaxDrainRetries:      maxDrainRetries,
			EvictGracePeriod:     evictGracePeriod,
			WaitBetweenRotations: waitBetweenRotations,
			WaitBetweenDrains:    waitBetweenDrains,
		}
		rotatorMetadata := &rotator.RotatorMetadata{}
		var err error
		rotatorMetadata, err = rotator.InitRotateCluster(&clusterRotator, rotatorMetadata, logger)
		if err != nil {
			return err
		}

		dryRun, _ := cmd.Flags().GetBool("dry-run")
		if dryRun {
			err := printJSON(rotatorMetadata)
			if err != nil {
				return errors.Wrap(err, "failed to print API request")
			}

			return nil
		}

		err = printJSON(rotatorMetadata)
		if err != nil {
			return errors.Wrap(err, "failed to print cluster response")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(rotateCmd)
	rotateCmd.Flags().String("cluster", "", "The name of the cluster to be upgraded.")
	rotateCmd.Flags().Int("max-scaling", 5, "The maximum number of nodes to rotate every time. If the number is bigger than the number of nodes, then the number of nodes will be the maximum number.")
	rotateCmd.Flags().Bool("rotate-masters", false, "if disabled, master nodes will not be rotated")
	rotateCmd.Flags().Bool("rotate-workers", true, "if disabled, worker nodes will not be rotated")
	rotateCmd.Flags().Int("max-drain-retries", 10, "The number of times to retry a node drain.")
	rotateCmd.Flags().Int("evict-grace-period", 600, "The pod eviction grace period when draining in seconds.")
	rotateCmd.Flags().Int("wait-between-rotations", 60, "Τhe time to wait between each rotation of a group of nodes.")
	rotateCmd.Flags().Int("wait-between-drains", 60, "The time to wait between each node drain in a group of nodes.")
	rotateCmd.MarkFlagRequired("cluster")
}
