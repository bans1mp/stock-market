import "./index.css";
import homeImage from "./assets/home.jpg";

function Navbar() {
  return (
    <nav className="navbar navbar-expand-lg navbar-dark" style={{ background: "linear-gradient(to bottom, #1e1e1e, #181818)" }}>
      <div className="container">
        <a className="navbar-brand text-white fw-bold fs-3" href="#">StockSim</a>
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
            <li className="nav-item">
              <a className="nav-link text-light" href="#">Home</a>
            </li>
            <li className="nav-item">
              <a className="nav-link text-light" href="#">Market</a>
            </li>
            <li className="nav-item">
              <a className="nav-link text-light" href="#">Contact</a>
            </li>
          </ul>
        </div>
      </div>
    </nav>
  );
}

function Footer() {
  return (
    <footer className="text-white text-center py-3 mt-auto" style={{ background: "linear-gradient(to top, #1e1e1e, #181818)" }}>
      <p className="mb-0">© 2025 StockSim | All rights reserved.</p>
    </footer>
  );
}

function App() {
  return (
    <div 
      className="d-flex flex-column min-vh-100 text-white" 
      style={{ 
        backgroundImage: `url(${homeImage})`, 
        backgroundSize: "cover", 
        backgroundPosition: "center", 
        backgroundRepeat: "no-repeat", 
        backgroundBlendMode: "overlay",
        backgroundColor: "rgba(0, 0, 0, 0.85)" 
      }}
    >
      <Navbar />
      <div className="container flex-grow-1 d-flex align-items-center justify-content-between mt-5 px-5">
        <div className="text-section" style={{ maxWidth: "50%", opacity: "0.75", color: "rgba(200,200,200,0.85)", textShadow: "2px 2px 6px rgba(0,0,0,0.8)" }}>
          <h1 className="fw-bold display-2 text-success border-start border-4 border-success ps-3" style={{ fontFamily: "Poppins, sans-serif" }}>
            Welcome to <br />
            <span className="text-danger">StockSim</span>
          </h1>
          <p className="lead fs-2">
            <span className="fw-bold text-success">Simulate</span>, <span className="fw-bold text-success">Trade</span>, <br />
            and <span className="fw-bold text-success">Master the Market</span> <br />
            with Realistic Stock Exchange Models.
          </p>
        </div>
        <div className="button-section d-flex align-items-center gap-4">
          <button className="btn btn-success px-5 py-4 fs-3 fw-bold" style={{ minWidth: "220px", opacity: "0.75" }}>Login</button>
          <button className="btn btn-outline-danger px-5 py-4 fs-3 fw-bold" style={{ minWidth: "220px", opacity: "0.75" }}>Register</button>
        </div>
      </div>
      <Footer />
    </div>
  );
}

export default App;