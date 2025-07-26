import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/in-touch/")({
    component: RouteComponent,
});

/**
 * Renders the UI for the "/in-touch/" route.
 *
 * Displays a message indicating the current route.
 */
function RouteComponent() {
    return <div>Hello "/in-touch/"!</div>;
}
