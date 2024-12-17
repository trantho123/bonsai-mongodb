import React, { useState } from 'react'
import { useNavigate, useParams } from 'react-router-dom'
import { Avatar, Button, CssBaseline, TextField, Typography, InputAdornment } from '@mui/material'
import { Box, Container } from '@mui/system'
import { MdLockOutline } from 'react-icons/md'
import { RiEyeFill, RiEyeOffFill } from 'react-icons/ri';
import axios from 'axios'
import { toast } from 'react-toastify'
import CopyRight from '../../Components/CopyRight/CopyRight'

const AddNewPassword = () => {
    const [newPassword, setNewPassword] = useState('')
    const [showPassword, setShowPassword] = useState(false);
    const [loading, setLoading] = useState(false);
    const navigate = useNavigate()
    const { token } = useParams()

    const handleSubmit = async (e) => {
        e.preventDefault()

        // Validate password
        if (!newPassword) {
            toast.error("Please enter a new password", {
                theme: "colored",
                autoClose: 2000
            });
            return;
        }

        try {
            setLoading(true);
            const response = await axios.post(
                `${process.env.REACT_APP_RESET_PASSWORD}`,
                {
                    token: token,
                    newPassword: newPassword
                }
            );

            if (response.data.success) {
                toast.success("Password has been reset successfully", {
                    theme: "colored",
                    autoClose: 2000
                });
                navigate('/login');
            }
        } catch (error) {
            toast.error(error.response?.data?.error || "Failed to reset password", {
                theme: "colored",
                autoClose: 2000
            });
        } finally {
            setLoading(false);
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
                    Reset Password
                </Typography>
                <Box component="form" onSubmit={handleSubmit} noValidate sx={{ mt: 1 }}>
                    <TextField
                        margin="normal"
                        required
                        fullWidth
                        name="password"
                        label="New Password"
                        type={showPassword ? "text" : "password"}
                        id="password"
                        value={newPassword}
                        onChange={(e) => setNewPassword(e.target.value)}
                        InputProps={{
                            endAdornment: (
                                <InputAdornment 
                                    position="end" 
                                    onClick={() => setShowPassword(!showPassword)}
                                    sx={{cursor:'pointer'}}
                                >
                                    {showPassword ? <RiEyeFill /> : <RiEyeOffFill />}
                                </InputAdornment>
                            )
                        }}
                    />
                    <Button
                        type="submit"
                        fullWidth
                        variant="contained"
                        sx={{ mt: 3, mb: 2 }}
                        disabled={loading}
                    >
                        {loading ? 'Resetting...' : 'Reset Password'}
                    </Button>
                </Box>
            </Box>
            <CopyRight sx={{ mt: 8, mb: 4 }} />
        </Container>
    )
}

export default AddNewPassword