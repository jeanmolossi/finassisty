// Types for the PWA runtime
/// <reference types="vite-plugin-pwa/client" />
import { RouterProvider } from "@tanstack/react-router";
import { createRoot } from "react-dom/client";
import { router } from "@/router/routes";
import "./index.css";
import { registerSW } from "virtual:pwa-register";

createRoot(document.getElementById("root")!).render(
    <RouterProvider router={router} />,
);

registerSW({ immediate: true });
