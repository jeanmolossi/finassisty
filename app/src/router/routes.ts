import { createBrowserRouter } from "react-router";
import Home from "../pages/home";
import App from "../App";
import InTouch from "../InTouch";

async function importDefault(path: string) {
    return import(path).then((res) => res.default);
}

export const router = createBrowserRouter([
    {
        path: "/",
        Component: App,
        children: [
            { index: true, Component: Home },
            {
                path: "/about",
                lazy: {
                    Component: async () => importDefault("../pages/about"),
                    loader: async () => importDefault("../pages/about/loader"),
                },
            },
            {
                path: "/dashboard",
                lazy: {
                    Component: async () => importDefault("../Dashboard"),
                },
            },
            { path: "/in-touch", Component: InTouch },
        ],
    },
]);
