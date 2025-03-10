// src/pages/Home.js
import React, { useState, useEffect } from 'react';
import { fetchChallengers, fetchGrandmasters, fetchMasters } from '../services/api';

const Home = () => {
  const [activeTab, setActiveTab] = useState('challenger');
  const [challengers, setChallengers] = useState([]);
  const [grandmasters, setGrandmasters] = useState([]);
  const [masters, setMasters] = useState([]);
  const [loading, setLoading] = useState(false);

  // Load data for the default "challenger" tab when the page loads
  useEffect(() => {
    loadData('challenger');
  }, []);

  const loadData = async (tab) => {
    setLoading(true); // Show loading spinner
    try {
      let data = [];
      if (tab === 'challenger') {
        data = await fetchChallengers();
        setChallengers(data);
      }
      if (tab === 'grandmaster') {
        data = await fetchGrandmasters();
        setGrandmasters(data);
      }
      if (tab === 'master') {
        data = await fetchMasters();
        setMasters(data);
      }
    } catch (error) {
      console.error(`Failed to fetch ${tab}`, error);
    } finally {
      setLoading(false); // Hide spinner
    }
  };

  const handleTabClick = (tab) => {
    setActiveTab(tab);
    loadData(tab); // Load data every time the tab is clicked
  };

  const renderTabContent = () => {
    if (loading) {
      return <p className="loading">Loading...</p>; // Spinner during loading
    }

    const data =
      activeTab === 'challenger'
        ? challengers
        : activeTab === 'grandmaster'
        ? grandmasters
        : masters;

    return (
      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>League Points</th>
            <th>Wins</th>
            <th>Losses</th>
          </tr>
        </thead>
        <tbody>
          {data.map((player) => (
            <tr key={player.summonerId || player.summonerName}>
              <td>{player.summonerName}</td>
              <td>{player.leaguePoints}</td>
              <td>{player.wins}</td>
              <td>{player.losses}</td>
            </tr>
          ))}
        </tbody>
      </table>
    );
  };

  return (
    <div className="home">
      <h1>Leaderboard</h1>
      <div className="tabs">
        <button
          className={`tab ${activeTab === 'challenger' ? 'active' : ''}`}
          onClick={() => handleTabClick('challenger')}
        >
          Challenger
        </button>
        <button
          className={`tab ${activeTab === 'grandmaster' ? 'active' : ''}`}
          onClick={() => handleTabClick('grandmaster')}
        >
          Grandmaster
        </button>
        <button
          className={`tab ${activeTab === 'master' ? 'active' : ''}`}
          onClick={() => handleTabClick('master')}
        >
          Master
        </button>
      </div>
      <div className="league-list">{renderTabContent()}</div>
    </div>
  );
};

export default Home;
