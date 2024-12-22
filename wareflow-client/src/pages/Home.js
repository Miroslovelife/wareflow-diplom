import React from 'react';
import { useAuth } from '../context/authContext';

const Home = () => {
    const { token } = useAuth();

    return (
        <div className="home-page">
            <h1>Welcome to Home Page</h1>
            {token && <p>You are logged in!</p>}
            <button onClick={() => alert('This is a protected route.')}>
                Click me
            </button>
        </div>
    );
};

export default Home;
