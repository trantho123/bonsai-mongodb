import React, { useEffect, useState } from 'react'
import { AiOutlineSearch } from 'react-icons/ai';
import {
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow,
    Paper,
    Container,
    InputAdornment,
    TextField,
    Chip,
    Box,
    Typography
} from '@mui/material'
import { Link } from 'react-router-dom';
import AddProduct from '../AddProduct';

const ProductTable = ({ data, getProductInfo }) => {
    const [filteredData, setFilteredData] = useState([]);
    const [searchTerm, setSearchTerm] = useState('');
    
    const columns = [
        {
            id: 'name',
            label: 'Name',
            minWidth: 170,
            align: 'center',
        },
        {
            id: 'image',
            label: 'Image',
            minWidth: 100,
            align: 'center',
        },
        {
            id: 'tags',
            label: 'Tags',
            align: 'center',
            minWidth: 150
        },
        {
            id: 'quantity',
            label: 'Quantity',
            align: 'center',
            minWidth: 100
        },
        {
            id: 'price',
            label: 'Price',
            minWidth: 100,
            align: 'center',
        },
        {
            id: 'rating',
            label: 'Rating',
            minWidth: 80,
            align: 'center',
        },
    ];

    const filterData = () => {
        if (!searchTerm) return data;
        
        const searchLower = searchTerm.toLowerCase();
        return data.filter((item) =>
            item.Name?.toLowerCase().includes(searchLower) ||
            item.description?.toLowerCase().includes(searchLower) ||
            item.tags?.some(tag => tag.name?.toLowerCase().includes(searchLower)) ||
            item.price?.toString().includes(searchLower) ||
            item.quantity?.toString().includes(searchLower) ||
            item.rating?.toString().includes(searchLower)
        );
    };

    const handleSearch = (event) => {
        setSearchTerm(event.target.value);
    };

    const formatCurrency = (amount) => {
        if (!amount && amount !== 0) return 'N/A';
        return new Intl.NumberFormat('vi-VN', { 
            style: 'currency', 
            currency: 'VND' 
        }).format(amount);
    };

    const getTagColor = (tagName) => {
        if (!tagName) return 'default';
        const colors = {
            book: 'primary',
            electronics: 'secondary',
            clothing: 'success',
            shoes: 'info',
            accessories: 'warning',
            default: 'default'
        };
        return colors[tagName.toLowerCase()] || colors.default;
    };

    useEffect(() => {
        setFilteredData(filterData());
    }, [data, searchTerm]);

    if (!data || data.length === 0) {
        return (
            <Box sx={{ textAlign: 'center', mt: 3 }}>
                <Typography variant="h6">No products available</Typography>
                <AddProduct getProductInfo={getProductInfo} data={[]} />
            </Box>
        );
    }

    return (
        <>
            <Container sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', marginBottom: 5, marginTop: 5 }}>
                <TextField
                    id="search"
                    type="search"
                    label="Search Products"
                    value={searchTerm}
                    onChange={handleSearch}
                    sx={{ width: { xs: 350, sm: 500, md: 800 } }}
                    InputProps={{
                        endAdornment: (
                            <InputAdornment position="end">
                                <AiOutlineSearch />
                            </InputAdornment>
                        ),
                    }}
                />
            </Container>
            
            <AddProduct getProductInfo={getProductInfo} data={data} />
            
            <Paper style={{ overflow: "auto", maxHeight: "500px" }}>
                <TableContainer sx={{ maxHeight: '500px' }}>
                    <Table stickyHeader aria-label="sticky table">
                        <TableHead>
                            <TableRow>
                                {columns.map((column) => (
                                    <TableCell
                                        key={column.id}
                                        align={column.align}
                                        style={{ 
                                            minWidth: column.minWidth, 
                                            color: "#1976d2",
                                            fontWeight: 'bold',
                                            backgroundColor: '#fff'
                                        }}
                                    >
                                        {column.label}
                                    </TableCell>
                                ))}
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {filteredData.length === 0 ? (
                                <TableRow>
                                    <TableCell colSpan={columns.length}>
                                        <div style={{ display: "flex", justifyContent: "center" }}>
                                            <h4>No products found.</h4>
                                        </div>
                                    </TableCell>
                                </TableRow>
                            ) : (
                                filteredData.map((product) => (
                                    <TableRow
                                        key={product._id}
                                        hover
                                        role="checkbox"
                                        tabIndex={-1}
                                    >
                                        <TableCell align="center">
                                            <Link to={`/admin/home/product/${product._id}`}>
                                                {product.name || 'Unnamed Product'}
                                            </Link>
                                        </TableCell>
                                        <TableCell align="center">
                                            <img 
                                                src={product.image} 
                                                alt={product.name || 'Product Image'} 
                                                style={{ 
                                                    width: "80px", 
                                                    height: "80px", 
                                                    objectFit: "contain" 
                                                }} 
                                            />
                                        </TableCell>
                                        <TableCell align="center">
                                            <Box sx={{ display: 'flex', gap: 0.5, flexWrap: 'wrap', justifyContent: 'center' }}>
                                                {product.tags && Array.isArray(product.tags) && product.tags.map((tag) => (
                                                    <Chip 
                                                        key={tag.id || tag._id}
                                                        label={tag.name || tag.Name || 'Unknown'}
                                                        color={getTagColor(tag.name || tag.Name)}
                                                        size="small"
                                                        sx={{ margin: '2px' }}
                                                    />
                                                ))}
                                            </Box>
                                        </TableCell>
                                        <TableCell align="center">
                                            {product.quantity || 0}
                                        </TableCell>
                                        <TableCell align="center">
                                            {formatCurrency(product.price)}
                                        </TableCell>
                                        <TableCell align="center">
                                            <Chip 
                                                label={product.rating ? product.rating.toFixed(1) : '0.0'}
                                                color={product.rating >= 4 ? 'success' : 'warning'}
                                                size="small"
                                            />
                                        </TableCell>
                                    </TableRow>
                                ))
                            )}
                        </TableBody>
                    </Table>
                </TableContainer>
            </Paper>
        </>
    );
}

export default ProductTable;