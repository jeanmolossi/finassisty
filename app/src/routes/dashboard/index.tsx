import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/dashboard/")({
    component: RouteComponent,
});

/**
 * Renders the greeting message for the "/dashboard/" route.
 *
 * Displays a simple message indicating the user is on the dashboard page.
 */
function RouteComponent() {
    return <div>Hello "/dashboard/"!</div>;
}
