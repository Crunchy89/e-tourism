import React from 'react';
import {StyleSheet} from "react-native"
import LinearGradient from 'react-native-linear-gradient';

const Container = (props:any) => {
    return (
        <LinearGradient start={{x: 0, y: 0}} end={{x: 1, y: 0}} colors={['rgba(255,226,89,.8)', 'rgba(255,204,86,.8)', 'rgba(255,167,81,.8)']} style={styles.news}>
        {props.children}
        </LinearGradient>
    );
};

const styles = StyleSheet.create({
    news:{
        borderRadius:5,
        margin:10,
        flex:3,
        overflow:"hidden",
        padding: 10,
    },
  });

export default Container;