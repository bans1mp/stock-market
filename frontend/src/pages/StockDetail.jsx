import axios from "axios";
import React, { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import Navbar from "../components/Navbar.jsx";
import Footer from "../components/Footer.jsx";
import { Line } from "react-chartjs-2";
import "chart.js/auto";
import homeImage from "../assets/home.jpg";

const StockDetail = () => {
  const { stockSymbol } = useParams();
  const navigate = useNavigate();
  const [historicalData, setHistoricalData] = useState([]);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchStockPrices = async () => {
      try {
        const token = localStorage.getItem("token");
        if (!token) {
          setError("Unauthorized: No token found");
          return;
        }

        const response = await axios.get(`http://localhost:8080/get-stock-prices/${stockSymbol}`, {
          headers: { Authorization: token },
        });
        
        setHistoricalData(response.data);
      } catch (err) {
        console.error("Error fetching stock prices:", err);
        setError("Failed to fetch stock prices");
      }
    };

    fetchStockPrices();
  }, [stockSymbol]);

  const chartData = {
    labels: historicalData.map(() => ""),
    datasets: [
      {
        label: `${stockSymbol} Price`,
        data: historicalData.map((data) => data.price),
        borderColor: "#28a745",
        backgroundColor: "rgba(40, 167, 69, 0.2)",
        borderWidth: 2,
        pointBackgroundColor: "#fff",
        pointBorderColor: "#28a745",
        tension: 0.4,
      },
    ],
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
          {stockSymbol}
        </h1>
        {error ? (
          <p className="fs-3 text-danger">{error}</p>
        ) : (
          <div className="w-75 bg-dark p-4 rounded shadow-lg mt-4" style={{ height: "400px" }}>
            <Line 
              data={chartData} 
              options={{ 
                responsive: true, 
                maintainAspectRatio: false,
                scales: {
                  x: {
                    display: false, // Hides X-axis labels
                  }
                }
              }} 
            />
          </div>
        )}
        <button
          className="btn btn-outline-light mt-4 px-4 py-2 fs-5 rounded-pill"
          onClick={() => navigate("/stocks")}
        >
          Back to Stocks
        </button>
      </div>
      <Footer />
    </div>
  );
};

export default StockDetail;
