import React, { useState } from 'react'
import { Avatar, Button, CssBaseline, TextField, Typography } from '@mui/material'
import { Box, Container } from '@mui/system'
import { MdLockOutline, MdMailOutline } from 'react-icons/md'
import axios from 'axios'
import { toast } from 'react-toastify'
import CopyRight from '../../Components/CopyRight/CopyRight'


const ForgotPasswordForm = () => {
    const [email, setEmail] = useState('')
    const [loading, setLoading] = useState(false)

    const handleSubmit = async (e) => {
        e.preventDefault()
        
        // Validate email
        const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/
        if (!emailRegex.test(email)) {
            toast.error("Please enter a valid email address", {
                theme: "colored",
                autoClose: 2000
            })
            return
        }

        try {
            setLoading(true)
            const response = await axios.post(
                `${process.env.REACT_APP_FORGOT_PASSWORD}`,
                { email }
            )

            if (response.data.success) {
                toast.success("Reset password link has been sent to your email", {
                    theme: "colored",
                    autoClose: 3000
                })
                setEmail('')
            }
        } catch (error) {
            toast.error(error.response?.data?.error || "Failed to process request", {
                theme: "colored",
                autoClose: 2000
            })
        } finally {
            setLoading(false)
        }
    }

    return (
        <Container component="main" maxWidth="xs">
            <CssBaseline />
            <Box
                sx={{
                    marginTop: 8,
                    display: 'flex',
                    flexDirection: 'column',
                    alignItems: 'center',
                }}
            >
                <Avatar sx={{ m: 1, bgcolor: '#1976d2' }}>
                    <MdLockOutline />
                </Avatar>
                <Typography component="h1" variant="h5">
                    Forgot Password
                </Typography>
                <Box component="form" onSubmit={handleSubmit} noValidate sx={{ mt: 1 }}>
                    <TextField
                        margin="normal"
                        required
                        fullWidth
                        id="email"
                        label="Email Address"
                        name="email"
                        autoComplete="email"
                        autoFocus
                        value={email}
                        onChange={(e) => setEmail(e.target.value)}
                    />
                    <Button
                        type="submit"
                        fullWidth
                        variant="contained"
                        sx={{ mt: 3, mb: 2 }}
                        disabled={loading}
                    >
                        {loading ? 'Sending...' : 'Send Reset Link'}
                    </Button>
                </Box>
            </Box>
            <CopyRight sx={{ mt: 8, mb: 4 }} />
        </Container>
    )
}

export default ForgotPasswordForm