import { Outlet, Link } from "react-router-dom";
import "./Content.css";

export default function Content() {
  return (
    <div className="Content">
      <nav className="Nav">
        <Link to="/">Home</Link>
        <Link to="/projects">Projects</Link>
        <Link to="/blog">Blog</Link>
        <Link to="/Resume.pdf" target={"_blank"}>
          Resume
        </Link>
      </nav>
      <hr></hr>
      <Outlet />
    </div>
  );
}
