import React, { Suspense } from "react";
import { BrowserRouter, Routes, Route, Link } from "react-router-dom";

const Home = React.lazy(() => import("./Home"));
const About = React.lazy(() => import("./About"));
const Dashboard = React.lazy(() => import("./Dashboard"));
const InTouch = React.lazy(() => import("./InTouch"));

export default function App() {
    return (
        <BrowserRouter>
            <nav className="p-4 flex gap-4 bg-gray-100">
                <Link to="/">Home</Link>
                <Link to="/about">About</Link>
                <Link to="/dashboard">Dashboard</Link>
                <Link to="/in-touch">In Touch</Link>
            </nav>
            <Routes>
                <Route
                    path="/"
                    element={
                        <Suspense fallback="Loading...">
                            <Home />
                        </Suspense>
                    }
                />
                <Route path="/about" element={<About />} />
                <Route path="/dashboard" element={<Dashboard />} />
                <Route path="/in-touch" element={<InTouch />} />
            </Routes>
        </BrowserRouter>
    );
}
