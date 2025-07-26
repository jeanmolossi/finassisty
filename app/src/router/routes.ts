import { createRouter, createRoute, createRootRoute, lazyRouteComponent } from "@tanstack/react-router";
import App from "@/App";
import Home from "@/pages/home";
import InTouch from "@/InTouch";

const rootRoute = createRootRoute({ component: App });

const indexRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: "/",
    component: Home,
});

export const aboutRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: "about",
    loader: () => import("@/pages/about/loader").then((m) => m.default()),
    component: lazyRouteComponent(() => import("@/pages/about")),
});

const dashboardRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: "dashboard",
    component: lazyRouteComponent(() => import("@/Dashboard")),
});

const inTouchRoute = createRoute({
    getParentRoute: () => rootRoute,
    path: "in-touch",
    component: InTouch,
});

const routeTree = rootRoute.addChildren([
    indexRoute,
    aboutRoute,
    dashboardRoute,
    inTouchRoute,
]);

export const router = createRouter({ routeTree });

declare module "@tanstack/react-router" {
    interface Register {
        router: typeof router;
    }
}
