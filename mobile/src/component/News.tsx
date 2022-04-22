import React from 'react';
import LinearGradient from 'react-native-linear-gradient';
import {Text,View,FlatList,StyleSheet} from "react-native";
import CarouselCards from './CaouselCard';

const News = () => {
    const [index, setIndex] = React.useState(0)
    const isCarousel = React.useRef(null)
    return (
        <LinearGradient start={{x: 0, y: 0}} end={{x: 1, y: 0}} colors={['rgba(255,226,89,.8)', 'rgba(255,204,86,.8)', 'rgba(255,167,81,.8)']} style={styles.news}>
        <Text style={styles.font}>NEWS</Text>
        <CarouselCards/>
        </LinearGradient>
    );
};

const styles = StyleSheet.create({
    font:{
        fontSize:15,
        marginBottom:5,
    },
    news:{
        borderRadius:5,
        marginTop:5,
        marginBottom:10,
        flex:2,
        justifyContent:"center",
        alignItems:"center",
        overflow:"hidden",
        padding: 5,
    },
  });

export default News;