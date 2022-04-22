import React from 'react';
import {Text,StyleSheet} from 'react-native';

const Header = (props:any) => {
    return (
        <Text style={styles.header}>{props.title}</Text>
    );
};

const styles = StyleSheet.create({
    header: {
        fontSize: 15,
        fontWeight: 'bold',
        marginBottom: 5,
    }
});

export default Header;