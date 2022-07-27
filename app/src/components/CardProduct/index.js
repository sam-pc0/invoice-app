import * as React from 'react';
import Typography from '@mui/material/Typography';
import Grid from '@mui/material/Grid';
import Card from '@mui/material/Card';
import CardActionArea from '@mui/material/CardActionArea';
import CardContent from '@mui/material/CardContent';


export default function CardProduct(props) {
    const { post } = props;
    return (

        <CardActionArea component="a" href="#">
            <Card style={{ marginLeft: '20px', marginRight: '20px', height: '180px', marginBottom: '40px' }} sx={{ display: 'flex' }}>
                <CardContent sx={{ flex: 1 }}>
                    <Typography component="h2" variant="h5">
                        {post.title}
                    </Typography>
                    <Typography variant="subtitle1" color="text.secondary">
                        {post.date}
                    </Typography>
                    <Typography variant="subtitle1" paragraph>
                        {post.description}
                    </Typography>

                </CardContent>
                <img
                    src={post.image}
                    alt={post.imageLabel}
                />
            </Card>
        </CardActionArea>

    )
}