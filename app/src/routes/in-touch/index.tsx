import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/in-touch/")({
    component: RouteComponent,
});

function RouteComponent() {
    return <div>Hello "/in-touch/"!</div>;
}
