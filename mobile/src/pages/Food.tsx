import React from 'react';
import {View,Text} from 'react-native';
import Container from '../component/Container';
import ButtonBack from '../component/ButtonBack';

const Food = ({navigation}) => {
    return (
        <>
        <ButtonBack link="Home" navigation={navigation} title="Food"/>
        <Container>
            <View>
                <Text>Food</Text>
            </View>
        </Container>
        </>
    );
};

export default Food;