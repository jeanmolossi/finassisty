import { StartServer } from "@tanstack/react-router-server/dist/esm/server";
import { createRouter } from "@tanstack/react-router";
import { routeTree } from "./routeTree.gen";

export default function entryServer() {
    const router = createRouter({
        routeTree,
        defaultPreload: "intent",
    }) as any; // eslint-disable-line @typescript-eslint/no-explicit-any
    return <StartServer router={router} />;
}
