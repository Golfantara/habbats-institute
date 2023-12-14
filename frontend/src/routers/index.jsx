import { Suspense } from "react";
import { lazily } from "react-lazily";
import { createBrowserRouter } from "react-router-dom";
import LoadingIcons from "react-loading-icons";
const { Register, Login, Home } = lazily(() => import("../components/index"));

const router = createBrowserRouter([
  {
    path: "/",
    element: (
      <Suspense fallback={<LoadingIcons.Bars />}>
        <Home />
      </Suspense>
    ),
  },

  {
    path: "/register",
    element: (
      <Suspense fallback={<LoadingIcons.Bars />}>
        <Register />
      </Suspense>
    ),
  },

  {
    path: "/login",
    element: (
      <Suspense fallback={<LoadingIcons.Bars />}>
        <Login />
      </Suspense>
    ),
  },
]);

export default router;
