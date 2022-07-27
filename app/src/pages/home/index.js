import * as React from 'react';
import AppBarComponent from '../../components/AppBar';
import HomaImage from '../../components/HomeImage';
import CardProduct from '../../components/CardProduct';
import Grid from '@mui/material/Grid';

export default function Home() {
    const featuredPosts = [
        {
            title: 'Logitech MK345 combo inalámbrico',
            date: 'US$39.99',
            description:
                'Teclado de computadora de tamaño completo con reposamanos. ',
            image: 'https://m.media-amazon.com/images/I/61S0sV1a57L._AC_SL1500_.jpg',
            imageLabel: 'Image Text',
        },
        {
            title: 'Logitech Cámara web C920S HD Pro',
            date: 'US$58.49',
            description:
                'Videollamadas Full HD 1080P: calidad de vídeo premium que te hace parecer un profesional.',
            image: 'https://m.media-amazon.com/images/I/61-6uAf8soL._AC_SL1500_.jpg',
            imageLabel: 'Image Text',
        },
        {
            title: 'Hamilton Beach Cafetera programable de 12 tazas',
            date: 'US$34.99',
            description:
                'Fácil de llenar: el depósito de agua es fácil de llenar y cuenta con una ventana.',
            image: 'https://m.media-amazon.com/images/I/81SmHZvpAkL._AC_SL1500_.jpg',
            imageLabel: 'Image Text',
        },
        {
            title: 'Juego de vajilla Gibson Home',
            date: 'US$55.59',
            description:
                'Soho Lounge: diseño contemporáneo que ejemplifica valor y calidad.',
            image: 'https://m.media-amazon.com/images/I/71G1vQOBcvL._AC_SL1500_.jpg',
            imageLabel: 'Image Text',
        },

        {
            title: 'Acer Aspire 5',
            date: 'US$480.52',
            description:
                'Pantalla panorámica IPS Full HD de 15.6 pulgadas (1920 x 1080) con retroiluminación LED ',
            image: 'https://m.media-amazon.com/images/I/71SCvh0L3OL._AC_SL1500_.jpg',
            imageLabel: 'Image Text',
        },

        {
            title: 'SAMSUNG QLED 4K UHD',
            date: 'US$1,897.99',
            description:
                'TV inteligente con múltiples asistentes de voz.',
            image: 'https://m.media-amazon.com/images/I/81inMIbTO9L._AC_SL1500_.jpg',
            imageLabel: 'Image Text',
        },
    ];

    return (
        <>
            <AppBarComponent />
            <HomaImage />
            <Grid container spacing={4}>

                {featuredPosts.map((post) => (
                    <Grid item xs={12} md={4}>
                        <CardProduct key={post.title} post={post} />
                    </Grid>
                ))}
            </Grid>
        </>
    );
}