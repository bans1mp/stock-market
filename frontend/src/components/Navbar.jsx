import { useEffect, useState } from "react";
import { Link } from "react-router-dom";

function Navbar() {
  const [token, setToken] = useState(localStorage.getItem("token"));

  useEffect(() => {
    const handleStorageChange = () => setToken(localStorage.getItem("token"));
    window.addEventListener("storage", handleStorageChange);
    return () => window.removeEventListener("storage", handleStorageChange);
  }, []);

  return (
    <nav className="navbar navbar-expand-lg navbar-dark" style={{ background: "linear-gradient(to bottom, #1e1e1e, #181818)" }}>
      <div className="container">
        <Link className="navbar-brand text-white fw-bold fs-3" to="/">StockSim</Link>
        <button
          className="navbar-toggler"
          type="button"
          data-bs-toggle="collapse"
          data-bs-target="#navbarNav"
        >
          <span className="navbar-toggler-icon"></span>
        </button>
        <div className="collapse navbar-collapse" id="navbarNav">
          <ul className="navbar-nav ms-auto fs-5">
            {token ? (
              <>
                <li className="nav-item">
                  <Link className="nav-link text-light" to="/">Home</Link>
                </li>
                <li className="nav-item">
                  <Link className="nav-link text-light" to="/stocks">Market</Link>
                </li>
              </>
            ) : (
              <>
                <li className="nav-item">
                  <Link className="nav-link text-light" to="/login">Login</Link>
                </li>
                <li className="nav-item">
                  <Link className="nav-link text-light" to="/signup">Signup</Link>
                </li>
              </>
            )}
          </ul>
        </div>
      </div>
    </nav>
  );
}

export default Navbar;
