package clients

import (
    "context"
    v1 "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"
    "github.com/solo-io/gloo/projects/gloo/pkg/bootstrap/clients/vault"
    "github.com/solo-io/solo-kit/pkg/api/v1/clients/factory"

    "github.com/hashicorp/vault/api"
    _ "github.com/hashicorp/vault/api/auth/aws"
)

type vaultSecretClientSettings struct {
    vault *api.Client

    // Vault's path where resources are located.
    root string

    // Tells Vault which secrets engine it should route traffic to. Defaults to "secret".
    // https://learn.hashicorp.com/tutorials/vault/getting-started-secrets-engines
    pathPrefix string
}

// The DefaultPathPrefix may be overridden to allow for non-standard vault mount paths
const DefaultPathPrefix = "secret"

type VaultClientInitFunc func() *api.Client

func NoopVaultClientInitFunc(c *api.Client) VaultClientInitFunc {
    return func() *api.Client {
        return c
    }
}

// NewVaultSecretClientFactory consumes a vault client along with a set of basic configurations for retrieving info with the client
func NewVaultSecretClientFactory(clientInit VaultClientInitFunc, pathPrefix, rootKey string) factory.ResourceClientFactory {
    return &factory.VaultSecretClientFactory{
        Vault:      clientInit(),
        RootKey:    rootKey,
        PathPrefix: pathPrefix,
    }
}

func VaultClientForSettings(vaultSettings *v1.Settings_VaultSecrets) (*api.Client, error) {
    return vault.NewAuthorizedClient(context.Background(), vaultSettings)
}
