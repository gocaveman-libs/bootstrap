package bootstrap

//go:generate mkwebresource -R -p "github.com/gocaveman-libs/bootstrap" -e "dist/.*/bootstrap[.](js|css)$" -r "github.com/gocaveman-libs/jquery,github.com/gocaveman-libs/popper.js" .

// vgo version uses semver imports: go:generate mkwebresource -R -p "github.com/gocaveman-libs/bootstrap/v4" -e "dist/.*/bootstrap[.](js|css)$" -r "github.com/gocaveman-libs/jquery/v3,github.com/gocaveman-libs/popper.js" .

