// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"fmt"
	"math/rand"
	"regexp"
	"testing"
	"time"

	tfe "github.com/hashicorp/go-tfe"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccTFETeamToken_basic(t *testing.T) {
	token := &tfe.TeamToken{}
	rInt := rand.New(rand.NewSource(time.Now().UnixNano())).Int()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccMuxedProviders,
		CheckDestroy:             testAccCheckTFETeamTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTFETeamToken_basic(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTFETeamTokenExists(
						"tfe_team_token.foobar", token),
				),
			},
		},
	})
}

func TestAccTFETeamToken_multiple_team_tokens(t *testing.T) {
	skipUnlessBeta(t)
	token := &tfe.TeamToken{}
	rInt := rand.New(rand.NewSource(time.Now().UnixNano())).Int()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccMuxedProviders,
		CheckDestroy:             testAccCheckTFETeamTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTFETeamToken_withMultipleTokens(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTFETeamTokenExists(
						"tfe_team_token.multi_token_1", token),
					testAccCheckTFETeamTokenExists(
						"tfe_team_token.multi_token_2", token),
					testAccCheckTFETeamTokenExists(
						"tfe_team_token.legacy", token),
				),
			},
		},
	})
}

func TestAccTFETeamToken_existsWithoutForce(t *testing.T) {
	token := &tfe.TeamToken{}
	rInt := rand.New(rand.NewSource(time.Now().UnixNano())).Int()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccMuxedProviders,
		CheckDestroy:             testAccCheckTFETeamTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTFETeamToken_basic(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTFETeamTokenExists(
						"tfe_team_token.foobar", token),
				),
			},

			{
				Config:      testAccTFETeamToken_existsWithoutForce(rInt),
				ExpectError: regexp.MustCompile(`token already exists`),
			},
		},
	})
}

func TestAccTFETeamToken_existsWithForce(t *testing.T) {
	token := &tfe.TeamToken{}
	rInt := rand.New(rand.NewSource(time.Now().UnixNano())).Int()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccMuxedProviders,
		CheckDestroy:             testAccCheckTFETeamTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTFETeamToken_basic(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTFETeamTokenExists(
						"tfe_team_token.foobar", token),
				),
			},

			{
				Config: testAccTFETeamToken_existsWithForce(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTFETeamTokenExists(
						"tfe_team_token.regenerated", token),
				),
			},
		},
	})
}

func TestAccTFETeamToken_invalidWithForceGenerateAndDescription(t *testing.T) {
	skipUnlessBeta(t)
	rInt := rand.New(rand.NewSource(time.Now().UnixNano())).Int()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccMuxedProviders,
		CheckDestroy:             testAccCheckTFETeamTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccTFETeamToken_WithForceGenerateAndDescription(rInt),
				ExpectError: regexp.MustCompile(`"force_regenerate" cannot be specified when "description"`),
			},
		},
	})
}

func TestAccTFETeamToken_withBlankExpiry(t *testing.T) {
	token := &tfe.TeamToken{}
	rInt := rand.New(rand.NewSource(time.Now().UnixNano())).Int()
	expiredAt := ""

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccMuxedProviders,
		CheckDestroy:             testAccCheckTFETeamTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTFETeamToken_withBlankExpiry(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTFETeamTokenExists(
						"tfe_team_token.foobar", token),
					resource.TestCheckResourceAttr(
						"tfe_team_token.foobar", "expired_at", expiredAt),
				),
			},
		},
	})
}

func TestAccTFETeamToken_withValidExpiry(t *testing.T) {
	token := &tfe.TeamToken{}
	rInt := rand.New(rand.NewSource(time.Now().UnixNano())).Int()
	expiredAt := "2051-04-11T23:15:59Z"

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccMuxedProviders,
		CheckDestroy:             testAccCheckTFETeamTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTFETeamToken_withValidExpiry(rInt),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTFETeamTokenExists(
						"tfe_team_token.expiry", token),
					resource.TestCheckResourceAttr(
						"tfe_team_token.expiry", "expired_at", expiredAt),
				),
			},
		},
	})
}

func TestAccTFETeamToken_withInvalidExpiry(t *testing.T) {
	rInt := rand.New(rand.NewSource(time.Now().UnixNano())).Int()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccMuxedProviders,
		CheckDestroy:             testAccCheckTFETeamTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccTFETeamToken_withInvalidExpiry(rInt),
				ExpectError: regexp.MustCompile(`must be a valid date or time, provided in iso8601 format`),
			},
		},
	})
}

func TestAccTFETeamToken_import(t *testing.T) {
	rInt := rand.New(rand.NewSource(time.Now().UnixNano())).Int()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccMuxedProviders,
		CheckDestroy:             testAccCheckTFETeamTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTFETeamToken_basic(rInt),
			},

			{
				ResourceName:            "tfe_team_token.foobar",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"token"},
			},
		},
	})
}

func TestAccTFETeamToken_importByTokenID(t *testing.T) {
	skipUnlessBeta(t)
	rInt := rand.New(rand.NewSource(time.Now().UnixNano())).Int()

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccMuxedProviders,
		CheckDestroy:             testAccCheckTFETeamTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTFETeamToken_withMultipleTokens(rInt),
			},
			{
				ResourceName:            "tfe_team_token.multi_token_1",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"token"},
			},
			{
				ResourceName:            "tfe_team_token.multi_token_2",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"token"},
			},
			{
				ResourceName:            "tfe_team_token.legacy",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"token"},
			},
		},
	})
}

func TestAccTFETeamToken_withNonexistentTeam(t *testing.T) {
	conf := `
resource "tfe_team_token" "invalid" {
  team_id    = "invalid"
}`
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV5ProviderFactories: testAccMuxedProviders,
		CheckDestroy:             testAccCheckTFETeamTokenDestroy,
		Steps: []resource.TestStep{
			{
				Config:      conf,
				ExpectError: regexp.MustCompile("resource not found, team does not exist or version of Terraform Enterprise\ndoes not support multiple team tokens with descriptions"),
			},
		},
	})
}

func testAccCheckTFETeamTokenExists(
	n string, token *tfe.TeamToken) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No instance ID is set")
		}

		var tt *tfe.TeamToken
		var err error
		if isTokenID(rs.Primary.ID) {
			tt, err = testAccConfiguredClient.Client.TeamTokens.ReadByID(ctx, rs.Primary.ID)
		} else {
			tt, err = testAccConfiguredClient.Client.TeamTokens.Read(ctx, rs.Primary.ID)
		}

		if err != nil {
			return err
		}

		if tt == nil {
			return fmt.Errorf("Team token not found")
		}

		*token = *tt

		return nil
	}
}

func testAccCheckTFETeamTokenDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "tfe_team_token" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No instance ID is set")
		}

		var err error
		if isTokenID(rs.Primary.ID) {
			_, err = testAccConfiguredClient.Client.TeamTokens.ReadByID(ctx, rs.Primary.ID)
		} else {
			_, err = testAccConfiguredClient.Client.TeamTokens.Read(ctx, rs.Primary.ID)
		}
		if err == nil {
			return fmt.Errorf("Team token %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

func testAccTFETeamToken_basic(rInt int) string {
	return fmt.Sprintf(`
resource "tfe_organization" "foobar" {
  name  = "tst-terraform-%d"
  email = "admin@company.com"
}

resource "tfe_team" "foobar" {
  name         = "team-test"
  organization = tfe_organization.foobar.id
}

resource "tfe_team_token" "foobar" {
  team_id = tfe_team.foobar.id
}`, rInt)
}

// NOTE: This config is invalid because you cannot manage multiple tokens for
// one team. It is expected to always error.
func testAccTFETeamToken_existsWithoutForce(rInt int) string {
	return fmt.Sprintf(`
resource "tfe_organization" "foobar" {
  name  = "tst-terraform-%d"
  email = "admin@company.com"
}

resource "tfe_team" "foobar" {
  name         = "team-test"
  organization = tfe_organization.foobar.id
}

resource "tfe_team_token" "foobar" {
  team_id = tfe_team.foobar.id
}

resource "tfe_team_token" "error" {
  team_id = tfe_team.foobar.id
}`, rInt)
}

// NOTE: This config is invalid because you cannot manage multiple tokens for
// one team. It can run without error _once_ due to the presence of
// force_regenerate, but is expected to error on any subsequent run.
func testAccTFETeamToken_existsWithForce(rInt int) string {
	return fmt.Sprintf(`
resource "tfe_organization" "foobar" {
  name  = "tst-terraform-%d"
  email = "admin@company.com"
}

resource "tfe_team" "foobar" {
  name         = "team-test"
  organization = tfe_organization.foobar.id
}

resource "tfe_team_token" "foobar" {
  team_id = tfe_team.foobar.id
}

resource "tfe_team_token" "regenerated" {
  team_id          = tfe_team.foobar.id
  force_regenerate = true
}`, rInt)
}

func testAccTFETeamToken_withBlankExpiry(rInt int) string {
	return fmt.Sprintf(`
resource "tfe_organization" "foobar" {
  name  = "tst-terraform-%d"
  email = "admin@company.com"
}

resource "tfe_team" "foobar" {
  name         = "team-test"
  organization = tfe_organization.foobar.id
}

resource "tfe_team_token" "foobar" {
  team_id = tfe_team.foobar.id
  expired_at = ""
}`, rInt)
}

func testAccTFETeamToken_withValidExpiry(rInt int) string {
	return fmt.Sprintf(`
resource "tfe_organization" "foobar" {
  name  = "tst-terraform-%d"
  email = "admin@company.com"
}

resource "tfe_team" "foobar" {
  name         = "team-test"
  organization = tfe_organization.foobar.id
}

resource "tfe_team_token" "expiry" {
  team_id    = tfe_team.foobar.id
  expired_at = "2051-04-11T23:15:59Z"
}`, rInt)
}

func testAccTFETeamToken_withInvalidExpiry(rInt int) string {
	return fmt.Sprintf(`
resource "tfe_organization" "foobar" {
  name  = "tst-terraform-%d"
  email = "admin@company.com"
}

resource "tfe_team" "foobar" {
  name         = "team-test"
  organization = tfe_organization.foobar.id
}

resource "tfe_team_token" "expiry" {
  team_id    = tfe_team.foobar.id
  expired_at = "2000-04-11"
}`, rInt)
}

func testAccTFETeamToken_withMultipleTokens(rInt int) string {
	return fmt.Sprintf(`
resource "tfe_organization" "foobar" {
  name  = "tst-terraform-%d"
  email = "admin@company.com"
}

resource "tfe_team" "foobar" {
  name         = "team-test"
  organization = tfe_organization.foobar.id
}


resource "tfe_team_token" "multi_token_1" {
  team_id     = tfe_team.foobar.id
  description = "tst-terraform-%d-token-1"
  expired_at  = "2051-04-11T23:15:59Z"
}

resource "tfe_team_token" "multi_token_2" {
  team_id    = tfe_team.foobar.id
  description = "tst-terraform-%d-token-2"
}

resource "tfe_team_token" "legacy" {
  team_id    = tfe_team.foobar.id
}`, rInt, rInt, rInt)
}

func testAccTFETeamToken_WithForceGenerateAndDescription(rInt int) string {
	return fmt.Sprintf(`
resource "tfe_organization" "foobar" {
  name  = "tst-terraform-%d"
  email = "admin@company.com"
}

resource "tfe_team" "foobar" {
  name         = "team-test"
  organization = tfe_organization.foobar.id
}


resource "tfe_team_token" "invalid" {
  team_id     = tfe_team.foobar.id
  description = "tst-terraform-%d-token"
  force_regenerate = true
}`, rInt, rInt)
}
