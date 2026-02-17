# Crossplane Provider for Cloud Foundry (DXFrontier Fork)

A fork of [`SAP/crossplane-provider-cloudfoundry`](https://github.com/SAP/crossplane-provider-cloudfoundry) (upstream `v0.3.3`) with Crossplane v2 support and custom IDP authentication.

---

## Installation

```yaml
apiVersion: pkg.crossplane.io/v1
kind: Provider
metadata:
  name: provider-cloudfoundry
spec:
  package: ghcr.io/dxfrontier/provider-cloudfoundry:<VERSION>
```

---

## Changes vs Upstream

### 1. Crossplane v2 / Namespace-scoped CRDs

Migrated from `crossplane-runtime v1` to `crossplane-runtime/v2 v2.1.0`. All 16 CRDs (14 managed resources + ProviderConfig + ProviderConfigUsage) are now **namespace-scoped** instead of cluster-scoped.

**What this means:**

- Every managed resource must include a `metadata.namespace`
- The `ProviderConfig` must be in the **same namespace** as the managed resources referencing it
- Cross-references between resources (e.g. Space → Organization) must be within the **same namespace**
- External Secret Store (ESS) support has been removed (dropped in Crossplane v2)
- Management Policies are always enabled (no feature flag needed)

**Breaking Change:** Existing cluster-scoped CRs must be deleted and recreated as namespaced resources.

### 2. Custom IDP Origin Support

Upstream omitted the `origin` parameter in CF UAA password grant flows, breaking authentication with custom Identity Providers. Fixed by including the IDP origin from ProviderConfig credentials.

---

## Releases

Published to GHCR under the `dxfrontier` org:
- **Provider package:** `ghcr.io/dxfrontier/provider-cloudfoundry`
- **Controller image:** `ghcr.io/dxfrontier/provider-cloudfoundry-controller`

---

**License:** Apache-2.0 | **Upstream:** [SAP/crossplane-provider-cloudfoundry](https://github.com/SAP/crossplane-provider-cloudfoundry)
