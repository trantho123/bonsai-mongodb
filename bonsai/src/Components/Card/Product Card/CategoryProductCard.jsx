import { 
    Card, 
    CardActionArea, 
    CardContent, 
    Rating, 
    Typography,
    Box
} from '@mui/material';
import { Link } from 'react-router-dom';
import styles from './ProductCard.module.css'

export default function CategoryProductCard({ product }) {
    console.log('CategoryProductCard received product:', product);
    
    if (!product) return null;

    return (
        <Card className={styles.main_card}>
            <Link 
                to={`/product/${product.ID || product.id}`} 
                style={{ textDecoration: 'none', color: 'inherit' }}
            >
                <CardActionArea>
                    <Box className={styles.cart_box}>
                        <img 
                            alt={product.Name || product.name || 'Product'} 
                            src={product.Image || product.image} 
                            loading='lazy' 
                            className={styles.cart_img} 
                        />
                    </Box>
                    <CardContent>
                        <Typography 
                            gutterBottom 
                            variant="h6" 
                            sx={{ textAlign: "center" }}
                        >
                            {(product.Name || product.name) ? 
                                ((product.Name || product.name).length > 20 
                                    ? (product.Name || product.name).slice(0, 20) + '...' 
                                    : (product.Name || product.name))
                                : 'Product Name'
                            }
                        </Typography>
                        
                        <Box sx={{ 
                            display: 'flex', 
                            justifyContent: 'space-between',
                            alignItems: 'center',
                            mt: 1
                        }}>
                            <Rating 
                                name="read-only" 
                                value={Number(product.Rating || product.rating) || 0}
                                precision={0.5}
                                readOnly 
                                size="small"
                            />
                            <Typography variant="h6" color="primary">
                                VND{product.Price || product.price || 0}
                            </Typography>
                        </Box>


                        {(product.Tags || product.tags) && (product.Tags || product.tags).length > 0 && (
                            <Box sx={{ mt: 1, display: 'flex', flexWrap: 'wrap', gap: 0.5 }}>
                                {(product.Tags || product.tags).map(tag => (
                                    <Typography 
                                        key={tag.ID || tag.id}
                                        variant="caption" 
                                        sx={{ 
                                            bgcolor: 'primary.light',
                                            color: 'white',
                                            px: 1,
                                            py: 0.25,
                                            borderRadius: 1
                                        }}
                                    >
                                        {tag.Name || tag.name}
                                    </Typography>
                                ))}
                            </Box>
                        )}
                    </CardContent>
                </CardActionArea>
            </Link>
        </Card>
    );
} 