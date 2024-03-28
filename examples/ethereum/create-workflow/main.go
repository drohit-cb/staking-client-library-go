/*
 * This example code, demonstrates creating an Ethereum Partial ETH (<32 ETH) workflow, the fundamental process handler for an E2E staking experience.
 */

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/coinbase/staking-client-library-go/auth"
	"github.com/coinbase/staking-client-library-go/client"
	"github.com/coinbase/staking-client-library-go/client/options"
	stakingpb "github.com/coinbase/staking-client-library-go/gen/go/coinbase/staking/orchestration/v1"
)

func main() {
	// TODO: Add your project ID found at cloud.coinbase.com or in your API key.
	projectID := ""

	ctx := context.Background()

	apiKey, err := auth.NewAPIKey(auth.WithLoadAPIKeyFromFile(true))
	if err != nil {
		log.Fatalf("error loading API key: %s", err.Error())
	}

	// Create a staking client.
	stakingClient, err := client.New(ctx, options.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("error instantiating staking client: %s", err.Error())
	}

	if projectID == "" {
		log.Fatalf("projectID must be set")
	}

	req := &stakingpb.CreateWorkflowRequest{
		Parent: fmt.Sprintf("projects/%s", projectID),
		Workflow: &stakingpb.Workflow{
			Action: "protocols/ethereum_kiln/networks/holesky/actions/stake",
			StakingParameters: &stakingpb.Workflow_EthereumKilnStakingParameters{
				EthereumKilnStakingParameters: &stakingpb.EthereumKilnStakingParameters{
					Parameters: &stakingpb.EthereumKilnStakingParameters_StakeParameters{
						StakeParameters: &stakingpb.EthereumKilnStakeParameters{
							StakerAddress:             "0xdb816889F2a7362EF242E5a717dfD5B38Ae849FE",
							IntegratorContractAddress: "0xA55416de5DE61A0AC1aa8970a280E04388B1dE4b",
							Amount: &stakingpb.Amount{
								Value:    "20",
								Currency: "ETH",
							},
						},
					},
				},
			},
			SkipBroadcast: true,
		},
	}

	workflow, err := stakingClient.Orchestration.CreateWorkflow(ctx, req)
	if err != nil {
		log.Fatalf("couldn't create workflow: %s", err.Error())
	}

	log.Printf("Workflow created: %s", workflow.Name)
}
