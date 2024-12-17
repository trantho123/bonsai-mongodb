import '../Login/login.css'
import React, { useEffect, useState } from 'react'
import axios from 'axios'
import { Link, useNavigate } from 'react-router-dom'
import { Avatar, Button, Checkbox, CssBaseline, FormControlLabel, Grid, InputAdornment, TextField, Typography } from '@mui/material'
import { MdLockOutline } from 'react-icons/md'
import { Box, Container } from '@mui/system'
import { toast } from 'react-toastify'
import CopyRight from '../../Components/CopyRight/CopyRight'
import { RiEyeFill, RiEyeOffFill } from 'react-icons/ri';



const Register = () => {

  const [credentials, setCredentials] = useState({ firstName: "", lastName: "", email: "", dob: "", password: "" })
  const [showPassword, setShowPassword] = useState(false);
  const navigate = useNavigate()
  const handleOnChange = (e) => {
    setCredentials({ ...credentials, [e.target.name]: e.target.value })
  }
  useEffect(() => {
    let auth = localStorage.getItem('Authorization');
    if (auth) {
      navigate("/")
    }
  }, [])
  const handleSubmit = async (e) => {
    e.preventDefault()
    let emailRegex = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    
    try {
        // Validate input fields
        if (!credentials.email || !credentials.firstName || !credentials.lastName || !credentials.password || !credentials.dob) {
            toast.error("All fields are required", { autoClose: 500, theme: 'colored' })
            return;
        }
        
        if (credentials.firstName.length < 3) {
            toast.error("Username must be at least 3 characters", { autoClose: 500, theme: 'colored' })
            return;
        }
        
        if (credentials.lastName.length < 3) {
            toast.error("Username must be at least 3 characters", { autoClose: 500, theme: 'colored' })
            return;
        }
        
        if (!emailRegex.test(credentials.email)) {
            toast.error("Please enter valid email", { autoClose: 500, theme: 'colored' })
            return;
        }
        
        if (credentials.password.length < 5) {
            toast.error("Password must be at least 5 characters", { autoClose: 500, theme: 'colored' })
            return;
        }

        // Send register request
        const sendAuth = await axios.post(`${process.env.REACT_APP_REGISTER}`, credentials)
        const receive = await sendAuth.data
        
        if (receive.success === true) {
            toast.success("Registered Successfully", { autoClose: 500, theme: 'colored' })
            localStorage.setItem('Authorization', receive.authToken)
            navigate('/')
        }
        
    } catch (error) {
        // Handle backend errors
        if (error.response) {
            // Backend validation errors
            if (error.response.data.error) {
                toast.error(error.response.data.error, { 
                    autoClose: 500, 
                    theme: 'colored' 
                })
            } 
            // Other backend errors
            else {
                toast.error("Registration failed. Please try again", { 
                    autoClose: 500, 
                    theme: 'colored' 
                })
            }
        } 
        // Network or other errors
        else {
            toast.error("Something went wrong. Please try again", { 
                autoClose: 500, 
                theme: 'colored' 
            })
        }
    }
  }


  return (
    <>
      <Container component="main" maxWidth="xs" sx={{ marginBottom: 10 }}>
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
            Sign up
          </Typography>
          <Box component="form" noValidate onSubmit={handleSubmit} sx={{ mt: 3 }}>
            <Grid container spacing={2}>
              <Grid item xs={12} sm={6}>
                <TextField
                  autoComplete="given-name"
                  name="firstName"
                  value={credentials.firstName}
                  onChange={handleOnChange}
                  required
                  fullWidth
                  id="firstName"
                  label="First Name"
                  autoFocus
                />
              </Grid>
              <Grid item xs={12} sm={6}>
                <TextField
                  required
                  fullWidth
                  id="lastName"
                  label="Last Name"
                  name="lastName"
                  value={credentials.lastName}
                  onChange={handleOnChange}
                  autoComplete="family-name"
                />
              </Grid>
              <Grid item xs={12}>
                <TextField
                  required
                  fullWidth
                  id="email"
                  label="Email Address"
                  name="email"
                  value={credentials.email}
                  onChange={handleOnChange}
                  autoComplete="email"
                />
              </Grid>
              <Grid item xs={12}>
                <TextField
                  required
                  fullWidth
                  id="dob"
                  label="Date of Birth"
                  name="dob"
                  type="date"
                  value={credentials.dob}
                  onChange={handleOnChange}
                  InputLabelProps={{
                    shrink: true,
                  }}
                />
              </Grid>
              <Grid item xs={12}>
                <TextField
                  required
                  fullWidth
                  name="password"
                  label="Password"
                  type={showPassword ? "text" : "password"}
                  id="password"
                  value={credentials.password}
                  onChange={handleOnChange}
                  InputProps={{
                    endAdornment: (
                      <InputAdornment position="end" onClick={() => setShowPassword(!showPassword)} sx={{ cursor: 'pointer' }}>
                        {showPassword ? <RiEyeFill /> : <RiEyeOffFill />}
                      </InputAdornment>
                    )
                  }}
                />
              </Grid>
            </Grid>
            <Button
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
            >
              Sign Up
            </Button>
            <Grid container justifyContent="flex-end">
              <Grid item>
                Already have an account?
                <Link to='/login' style={{ color: '#1976d2', marginLeft: 3 }}>
                  Sign in
                </Link>
              </Grid>
            </Grid>
          </Box>
        </Box>
        <CopyRight sx={{ mt: 5 }} />
      </Container>
    </>
  )
}

export default Register