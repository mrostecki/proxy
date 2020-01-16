workspace(name = "cilium")

#
# We grep for the following line to generate SOURCE_VERSION file for non-git
# distribution builds. This line must start with the string ENVOY_SHA followed by
# an equals sign and a git SHA in double quotes.
#
# No other line in this file may have ENVOY_SHA followed by an equals sign!
#
ENVOY_SHA = "9153a6077d17ed4af1457b998a9a6b3c75572456"
ENVOY_SHA256 = "53467391b515ac088d7d89f8d177669ff5e3486f1c9d1d42fdab940e6bc0c04b"

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "envoy",
    url = "https://github.com/jrajahalme/envoy/archive/" + ENVOY_SHA + ".tar.gz",
    sha256 = ENVOY_SHA256,
    strip_prefix = "envoy-" + ENVOY_SHA,
    patch_args = ["-p1"],
)

#
# Bazel does not do transitive dependencies, so we must basically
# include all of Envoy's WORKSPACE file below, with the following
# changes:
# - Skip the 'workspace(name = "envoy")' line as we already defined
#   the workspace above.
# - loads of "//..." need to be renamed as "@envoy//..."
#

load("@envoy//bazel:api_binding.bzl", "envoy_api_binding")
envoy_api_binding()

load("@envoy//bazel:api_repositories.bzl", "envoy_api_dependencies")
envoy_api_dependencies()

load("@envoy//bazel:repositories.bzl", "envoy_dependencies")
envoy_dependencies()

load("@envoy//bazel:dependency_imports.bzl", "envoy_dependency_imports")
envoy_dependency_imports()

# Dependencies for Istio filters.
# Cf. https://github.com/istio/proxy.
# Version 1.2.2
# ISTIO_PROXY_SHA = "a975561b980463f08689d3debe33bb9eefc80c3d"
# ISTIO_PROXY_SHA256 = "c0123fe73be4c9f2fe5e673952743ceb836f5972a8377ea876d90b7ab63af6eb"

#http_archive(
#    name = "istio_proxy",
#    url = "https://github.com/istio/proxy/archive/" + ISTIO_PROXY_SHA + ".tar.gz",
#    sha256 = ISTIO_PROXY_SHA256,
#    strip_prefix = "proxy-" + ISTIO_PROXY_SHA,
#)

#load("@istio_proxy//:repositories.bzl", "mixerapi_dependencies")
#mixerapi_dependencies()

#bind(
#    name = "boringssl_crypto",
#    actual = "//external:ssl",
#)
