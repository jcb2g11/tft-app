// src/services/api.js
import axios from 'axios';

const API_URL = process.env.NODE_ENV === 'production' 
  ? 'https://tft-app-qjx6.onrender.com'  // Replace with actual API URL
  : 'http://localhost:8080';

export const fetchChallengers = async () => {
  try {
    const response = await axios.get(`${API_URL}/challenger`);
    return response.data;
  } catch (error) {
    console.error('Error fetching challengers:', error);
    return [];
  }
};

export const fetchGrandmasters = async () => {
  try {
    const response = await axios.get(`${API_URL}/grandmaster`);
    return response.data;
  } catch (error) {
    console.error('Error fetching grandmasters:', error);
    return [];
  }
};

export const fetchMasters = async () => {
  try {
    const response = await axios.get(`${API_URL}/master`);
    return response.data;
  } catch (error) {
    console.error('Error fetching masters:', error);
    return [];
  }
};
