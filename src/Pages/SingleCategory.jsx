import { useState, useEffect } from 'react';
import { Container, Grid, Typography, Box } from '@mui/material';
import CategoryProductCard from '../Components/Card/Product Card/CategoryProductCard';
import axios from 'axios';

export default function SingleCategory({ category }) {
    const [products, setProducts] = useState([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        fetchProducts();
    }, [category]);

    const fetchProducts = async () => {
        try {
            console.log('Fetching products for category:', category);
            const response = await axios.get(
                `${process.env.REACT_APP_FETCH_PRODUCT}/category/${category}`
            );
            console.log('Products response:', response.data);
            setProducts(response.data);
            setLoading(false);
        } catch (error) {
            console.error('Error fetching products:', error);
            setLoading(false);
        }
    };

    if (loading) return <div>Loading...</div>;

    return (
        <Container maxWidth="xl" sx={{ mt: 4, mb: 8 }}>
            <Typography variant="h4" sx={{ mb: 4, textAlign: 'center' }}>
                {category}
            </Typography>

            <Grid container spacing={3}>
                {products.map(product => (
                    <Grid item xs={12} sm={6} md={4} lg={3} key={product.id}>
                        <CategoryProductCard product={product} />
                    </Grid>
                ))}
            </Grid>

            {products.length === 0 && (
                <Box sx={{ textAlign: 'center', mt: 4 }}>
                    <Typography variant="h6" color="text.secondary">
                        No products found in this category
                    </Typography>
                </Box>
            )}
        </Container>
    );
} 