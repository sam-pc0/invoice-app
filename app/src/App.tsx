import React from 'react';
import logo from './logo.svg';
import { Routes, Route, Link } from "react-router-dom";
import Login from '../src/pages/login';
import Home from '../src/pages/home';
import MyProfile from './pages/MiPerfil';
import './App.css';

function App() {
  return (
    <Routes>
      <Route path="/" element={<Login />} />
      <Route path="/home" element={<Home />} />
      <Route path="/myProfile" element={<MyProfile />} />
    </Routes>

  );
}

export default App;
