package main

import (
	"github.com/pulumi/pulumi-github/sdk/v6/go/github"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-vercel/sdk/go/vercel"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Replace with your repository owner and name
		repoOwner := "mojtabatorabi"
		repoName := "reactf"

		// Get information about the existing GitHub repository
		repo, err := github.LookupRepository(ctx, &github.LookupRepositoryArgs{
			//Owner: pulumi.String(repoOwner),
			//Name:     pulumi.String(repoName),
			FullName: pulumi.StringRef(repoOwner + "/" + repoName),
		})
		if err != nil {
			return err
		}

		// Create a Vercel project
		project, err := vercel.NewProject(ctx, "my-vercel-project", &vercel.ProjectArgs{
			Name: pulumi.String("my-vercel-project"),
			//  TeamId:      pulumi.String("<your-vercel-team-id>"), // replace with your Vercel team ID
			GitRepository: &vercel.ProjectGitRepositoryArgs{
				Repo:             pulumi.String(repo.Name),
				Type:             pulumi.String("github"),
				ProductionBranch: pulumi.String("master"),
			},
			Framework:      pulumi.String("create-react-app"), // replace with your framework
			BuildCommand:   pulumi.String("npm run build"),    // replace with your build command
			DevCommand:     pulumi.String("npm run dev"),      // replace with your dev command
			InstallCommand: pulumi.String("npm install"),      // replace with your install command
			//OutputDirectory:          pulumi.String("out"),              // replace with your output directory
			PublicSource:             pulumi.Bool(true),
			ServerlessFunctionRegion: pulumi.String("iad1"), // replace with your desired region
		})
		if err != nil {
			return err
		}

		// Create a Vercel deployment
		_, err = vercel.NewDeployment(ctx, "my-vercel-deployment", &vercel.DeploymentArgs{
			ProjectId: project.ID(),
			//   TeamId:         pulumi.String("<your-vercel-team-id>"), // replace with your Vercel team ID
			Production:      pulumi.Bool(true),
			DeleteOnDestroy: pulumi.Bool(true),
			Ref:             pulumi.String("master"),
		})
		if err != nil {
			return err
		}

		// Export the repository name and URL as stack outputs
		//ctx.Export("repositoryName", pulumi.String(repo.Name))
		//ctx.Export("repositoryUrl", pulumi.String(repo.HtmlUrl))

		return nil
	})
}
