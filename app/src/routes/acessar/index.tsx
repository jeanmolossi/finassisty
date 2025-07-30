import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/acessar/")({
    component: AcessarPage,
});

function AcessarPage() {
    return (
        <main className="p-4">
            <h1 className="text-2xl font-bold mb-4">Acessar</h1>
            <a href="/api/v1/auth/google/login" className="text-blue-600 underline">
                Entrar com Google
            </a>
        </main>
    );
}
