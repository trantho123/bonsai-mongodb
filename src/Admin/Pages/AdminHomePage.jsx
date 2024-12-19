import React, { useEffect, useState } from 'react'
import axios from 'axios'
import { useNavigate } from 'react-router-dom'
import { toast } from 'react-toastify';
import { Container, Alert, CircularProgress, Box } from '@mui/material';
import BasicTabs from '../Components/AdminTabs';
import CopyRight from '../../Components/CopyRight/CopyRight'

const AdminHomePage = () => {
    const [user, setUser] = useState([]);
    const [isAdmin, setAdmin] = useState(false);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        getUser();
    }, [])

    let navigate = useNavigate()
    let authToken = localStorage.getItem("Authorization")
    let userRole = localStorage.getItem("Role")

    const getUser = async () => {
        try {
            // Log auth details for debugging
            console.log("Auth Token:", authToken)
            console.log("User Role:", userRole)

            if (!authToken) {
                setError("No authentication token found. Please login again.")
                setLoading(false)
                return
            }

            const { data } = await axios.get(`${process.env.REACT_APP_ADMIN_GET_ALL_USERS}`, {
                headers: {
                    'Authorization': `Bearer ${authToken}`,
                }
            })
            
            setUser(data.data || [])
            setAdmin(true)
            setError(null)
        } catch (error) {
            console.error("API Error:", error.response || error)
            
            let errorMessage = "An unknown error occurred"
            if (error.response) {
                // Get detailed error message from API
                errorMessage = error.response.data.details || error.response.data.error || error.response.statusText
                
                // Handle specific error codes
                switch(error.response.status) {
                    case 401:
                        errorMessage = "Your session has expired. Please login again."
                        navigate('/admin/login')
                        break
                    case 403:
                        errorMessage = "You don't have permission to access this page. Admin privileges required."
                        break
                    case 500:
                        errorMessage = "Server error. Please try again later."
                        break
                }
            }
            
            setError(errorMessage)
            !isAdmin && navigate('/')
            toast.error(errorMessage, { 
                autoClose: 3000, 
                theme: "colored",
                position: "top-center"
            });
        } finally {
            setLoading(false)
        }
    }

    if (loading) {
        return (
            <Box display="flex" justifyContent="center" alignItems="center" minHeight="100vh">
                <CircularProgress />
            </Box>
        )
    }

    return (
        <>
            {error && (
                <Alert severity="error" sx={{ margin: 2 }}>
                    {error}
                </Alert>
            )}
            
            {isAdmin && (
                <Container maxWidth="100%">
                    <h1 style={{ textAlign: "center", margin: "20px 0", color: "#1976d2" }}>
                        Dashboard
                    </h1>
                    <BasicTabs user={user} getUser={getUser} />
                </Container>
            )}
            
            <CopyRight sx={{ mt: 8, mb: 10 }} />
        </>
    )
}

export default AdminHomePage