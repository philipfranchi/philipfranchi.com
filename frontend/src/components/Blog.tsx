import { Outlet } from "react-router-dom";

export default function Blog() {
  return (
    <main>
      <h2>Blog</h2>
      <Outlet />
    </main>
  );
}
