
import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import {AuthProvider} from './context/authContext';
import Home from './pages/Home';
import Login from './pages/Login';

const routes = () => (
    <Router>
        <Routes>
            <Route path="/login" element={<Login />} />
            <Route path="/" element={
                <AuthProvider>
                    <Home />
                </AuthProvider>
            } />
        </Routes>
    </Router>
);

export default routes;
