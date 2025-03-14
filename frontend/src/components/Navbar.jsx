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

export default Navbar