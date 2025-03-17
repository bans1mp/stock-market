import axios from "axios";
import React, { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import Navbar from "../components/Navbar.jsx";
import Footer from "../components/Footer.jsx";
import homeImage from "../assets/home.jpg";

const Stocks = () => {
  const [stocks, setStocks] = useState([]);
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchStocks = async () => {
      try {
        const token = localStorage.getItem("token");
        if (!token) {
          setError("Unauthorized: No token found");
          return;
        }

        const response = await axios.get("http://localhost:8080/get-stocks", {
          headers: {
            Authorization: token,
          },
        });

        setStocks(response.data);
      } catch (err) {
        console.error("Error fetching stocks:", err);
        setError("Failed to fetch stocks");
      }
    };

    fetchStocks();
  }, []);

  const handleOrder = (symbol, type) => {
    navigate(`/${type}-order/${symbol}`);
  };

  return (
    <div
      className="d-flex flex-column min-vh-100 text-white"
      style={{
        backgroundImage: `url(${homeImage})`,
        backgroundSize: "cover",
        backgroundPosition: "center",
        backgroundRepeat: "no-repeat",
        backgroundBlendMode: "overlay",
        backgroundColor: "rgba(0, 0, 0, 0.85)",
      }}
    >
      <Navbar />
      <div className="container flex-grow-1 d-flex flex-column align-items-center justify-content-center mt-5 px-5">
        <h1
          className="fw-bold display-2 text-success border-start border-4 border-success ps-3 mb-4"
          style={{ fontFamily: "Poppins, sans-serif" }}
        >
          Market <span className="text-danger">Stocks</span>
        </h1>
        {error ? (
          <p className="fs-3 text-danger">{error}</p>
        ) : (
          <div className="table-responsive w-75">
            <table className="table table-dark table-striped text-center">
              <thead>
                <tr>
                  <th>Stock Symbol</th>
                  <th>IPO Price ($)</th>
                  <th>Actions</th>
                </tr>
              </thead>
              <tbody>
                {stocks.length > 0 ? (
                  stocks.map((stock, index) => (
                    <tr key={index}>
                      <td>{stock.symbol}</td>
                      <td>{stock.ipo_price}</td>
                      <td>
                        <button
                          className="btn btn-success me-2"
                          onClick={() => handleOrder(stock.symbol, "buy")}
                        >
                          Buy
                        </button>
                        <button
                          className="btn btn-danger"
                          onClick={() => handleOrder(stock.symbol, "sell")}
                        >
                          Sell
                        </button>
                      </td>
                    </tr>
                  ))
                ) : (
                  <tr>
                    <td colSpan="3" className="text-center fs-4 py-3">
                      Loading stocks...
                    </td>
                  </tr>
                )}
              </tbody>
            </table>
          </div>
        )}
      </div>
      <Footer />
    </div>
  );
};

export default Stocks;
