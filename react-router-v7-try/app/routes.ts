import type { RouteConfig } from '@react-router/dev/routes'
import { index, layout, route } from '@react-router/dev/routes'

export const routes: RouteConfig = [
  index('routes/home.tsx'),

  
  route("dashboard", "routes/dashboard/layout.tsx", [
    index("routes/dashboard/home.tsx"),
  ]),
]
