import React, { useState, useEffect } from 'react'
import axios from 'axios'
import { Container, Box, Grid, TextField, InputAdornment, Typography } from '@mui/material'
import { BiSearch } from 'react-icons/bi'
import Loading from '../Components/loading/Loading'
import CategoryProductCard from '../Components/Card/Product Card/CategoryProductCard'

const SingleCategory = () => {
    const [products, setProducts] = useState([])
    const [isLoading, setIsLoading] = useState(false)
    const [searchTerm, setSearchTerm] = useState('')
    const [displayedProducts, setDisplayedProducts] = useState([])

    useEffect(() => {
        getAllProducts()
        window.scroll(0, 0)
    }, [])

    // Search sản phẩm chứa ký tự được nhập
    useEffect(() => {
        console.log('Search effect running with:', {
            searchTerm,
            productsLength: products.length,
            sampleProduct: products[0] // Xem cấu trúc của product
        });

        if (searchTerm.trim() === '') {
            setDisplayedProducts(products)
        } else {
            // Log để kiểm tra cấu trúc dữ liệu
            console.log('First product:', products[0]);
            
            const filtered = products.filter(product => {
                // Kiểm tra cấu trúc của từng product
                console.log('Checking product:', {
                    name: product.Name || product.name,
                    searchTerm: searchTerm,
                    includes: product.Name?.toLowerCase().includes(searchTerm.toLowerCase())
                });
                
                return product.Name?.toLowerCase().includes(searchTerm.toLowerCase()) ||
                       product.name?.toLowerCase().includes(searchTerm.toLowerCase())
            });

            console.log('Filtered results:', filtered);
            setDisplayedProducts(filtered)
        }
    }, [products, searchTerm])

    const getAllProducts = async () => {
        try {
            setIsLoading(true)
            const response = await axios.get(`${process.env.REACT_APP_FETCH_PRODUCT}`)
            // Log response để xem cấu trúc data
            console.log('API Response:', {
                status: response.status,
                data: response.data,
                sampleProduct: response.data[0]
            });
            
            setProducts(response.data)
            setDisplayedProducts(response.data)
            setIsLoading(false)
        } catch (error) {
            console.error('Error fetching products:', {
                message: error.message,
                response: error.response?.data
            })
            setIsLoading(false)
        }
    }

    const handleSearch = (e) => {
        const value = e.target.value
        console.log('Search term:', value) // Debug log
        setSearchTerm(value)
    }

    if (isLoading) {
        return (
            <Container maxWidth='xl' sx={{ mt: 10, display: "flex", justifyContent: "center", flexWrap: "wrap", pb: 2 }}>
                <Loading />
                <Loading />
                <Loading />
                <Loading />
            </Container>
        )
    }

    return (
        <Container maxWidth='xl' sx={{ mt: 9 }}>
            {/* Search Box */}
            <Box sx={{ mb: 4, display: 'flex', justifyContent: 'center' }}>
                <TextField
                    fullWidth
                    sx={{ maxWidth: 600 }}
                    placeholder="Enter product name to search..."
                    value={searchTerm}
                    onChange={handleSearch}
                    InputProps={{
                        startAdornment: (
                            <InputAdornment position="start">
                                <BiSearch />
                            </InputAdornment>
                        ),
                    }}
                />
            </Box>

            {/* Products Grid */}
            <Grid container spacing={3} justifyContent="center">
                {displayedProducts.map(product => (
                    <Grid item xs={12} sm={6} md={4} lg={3} key={product.ID}>
                        <CategoryProductCard product={product} />
                    </Grid>
                ))}
            </Grid>

            {/* No Results Message */}
            {searchTerm.trim() !== '' && displayedProducts.length === 0 && (
                <Box sx={{ textAlign: 'center', mt: 4 }}>
                    <Typography variant="h6" color="text.secondary">
                        No products found containing "{searchTerm}"
                    </Typography>
                </Box>
            )}
        </Container>
    )
}

export default SingleCategory

    //         