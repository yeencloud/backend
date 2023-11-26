package usecases

import (
	"back/src/core/domain"
	"github.com/davecgh/go-spew/spew"
	"github.com/rs/zerolog/log"
	"k8s.io/client-go/rest"
)

func (i interactor) steps() []domain.SetupStep {
	return []domain.SetupStep{
		{
			StepIdentifier: "SelectLanguage",
			VerifyDone: func() bool {
				return false
			},
			RequiredFields: func() []domain.SetupField {
				availableLanguages := i.translator.LanguageTags()
				fields := []string{}

				for _, lang := range availableLanguages {
					fields = append(fields, lang.String())
				}

				return []domain.SetupField{
					{
						Identifier:   "language",
						Name:         "Language",
						Type:         "select",
						SelectValues: fields,
					},
				}
			},
		},
		{
			StepIdentifier: "CreateAdminUser",
			VerifyDone: func() bool {
				return i.userRepo.CountUsers() > 0
			},
			RequiredFields: func() []domain.SetupField {
				return []domain.SetupField{
					{
						Identifier: "name",
						Name:       "name",
						Type:       "text",
					},
					{
						Identifier: "email",
						Name:       "Email",
						Type:       "email",
					},
					{
						Identifier: "password",
						Name:       "Password",
						Type:       "password",
					},
				}
			},
		},
		{
			StepIdentifier: "ImportKubernetesCluster",
			RequiredFields: func() []domain.SetupField {
				config, err := rest.InClusterConfig()
				strcfg := ""
				if err != nil {
					println(err.Error())
				} else {
					strcfg = config.String()
				}

				return []domain.SetupField{
					{
						Identifier: "kubeconfig",
						Name:       "Kubeconfig",
						Type:       "textarea",
						Value:      strcfg,
					},
				}
			},
		},
		{
			StepIdentifier: "AddS3Repository",
		},
		{
			StepIdentifier: "AddMutualizedPostgresServer",
		},
	}
}

func (i *interactor) prepareSetupSteps() {
	steps := i.steps()

	spew.Dump(steps[2].RequiredFields())

	for _, step := range steps {
		if step.VerifyDone == nil {
			log.Warn().Str("step", step.StepIdentifier).Msg("No verification function defined for this step, it will be skipped")
		} else {
			isDone := step.VerifyDone()
			if isDone {
				log.Info().Str("step", step.StepIdentifier).Msg("Step is already done")
			} else {
				log.Info().Str("step", step.StepIdentifier).Msg("Step is not done yet")
				i.remainingSetupSteps = append(i.remainingSetupSteps, step.StepIdentifier)
			}
		}
	}

	log.Info().Int("remainingSteps", len(i.remainingSetupSteps)).Strs("steps", i.remainingSetupSteps).Msg("Remaining setup steps")
}

func (i interactor) VerifySetup() {
	println("VerifySetup")
}

func (i interactor) IsSetupDone() (bool, string) {
	done := len(i.remainingSetupSteps) == 0
	if done {
		return true, ""
	}
	return false, i.remainingSetupSteps[0]
}
