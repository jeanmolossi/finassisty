import { createFileRoute, redirect } from "@tanstack/react-router";

export const Route = createFileRoute("/dashboard/")({
    beforeLoad: async () => {
        const res = await fetch("/api/v1/auth/session", {
            credentials: "include",
        });
        if (!res.ok) {
            throw redirect({ to: "/acessar" });
        }
    },
    component: RouteComponent,
});

function RouteComponent() {
    return <div>Hello "/dashboard/"!</div>;
}
