import axios from 'axios';
import { toast } from 'react-toastify';

const api = axios.create();

// Request interceptor
api.interceptors.request.use(
    config => {
        const token = localStorage.getItem('Authorization');
        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
    error => {
        return Promise.reject(error);
    }
);

// Response interceptor
api.interceptors.response.use(
    response => response,
    error => {
        if (error.response?.status === 401) {
            // Clear all auth data
            localStorage.clear();
            
            // Show message
            toast.error('Session expired. Please login again');
            
            // Redirect to appropriate login page
            const isAdminRoute = window.location.pathname.startsWith('/admin');
            if (isAdminRoute) {
                window.location.href = '/admin/login';
            } else {
                window.location.href = '/login';
            }
        }
        return Promise.reject(error);
    }
);

export default api; 