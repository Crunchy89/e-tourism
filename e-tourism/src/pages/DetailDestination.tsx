import React from 'react';
import {View,Text} from 'react-native';
import Container from '../component/Container';
import ButtonBack from '../component/ButtonBack';

const DetailDestination = ({route,navigation}) => {
    const { id, title } = route.params;
    return (
        <>
        <ButtonBack link="Destination" navigation={navigation} title={title}/>
        <Container>
            <View>
                <Text>{title}</Text>
            </View>
        </Container>
        </>
    );
};

export default DetailDestination;