import React from 'react';
import {View,Text,Image, StyleSheet, FlatList} from 'react-native';
import Container from '../component/Container';
import ButtonBack from '../component/ButtonBack';
import {data} from "../component/data"
import Header from '../component/Header';

const DetailDestination = ({route,navigation}) => {
    const { id, title } = route.params;
    return (
        <>
        <ButtonBack link="Destination" navigation={navigation} title={title}/>
        <Container>
            <View>
                <FlatList
                horizontal
                keyExtractor={(_,item) => item.toString()}
                data={data}
                renderItem={({item})=>(
                    <Image style={styles.image} source={{uri:item.imgUrl}} />
                )}
                showsHorizontalScrollIndicator={false}
                />
            </View>
            <Header title="Description" />
            <Text style={styles.text}>
                Lorem ipsum dolor, sit amet consectetur adipisicing elit. Reprehenderit natus libero ratione dolore ab eos commodi maxime earum doloribus sequi repudiandae ex eius modi repellendus consequatur, voluptate reiciendis cum quam.
                Lorem ipsum, dolor sit amet consectetur adipisicing elit. Minima quos veritatis eum quam modi quisquam dolorem aliquam, eius, exercitationem dolorum eos voluptate consectetur ea asperiores! Esse sint dolore nemo incidunt?
            </Text>
            <Header title="Facility" />
            <FlatList
                horizontal
                data={[]}
                keyExtractor={(_,i)=>i.toString()}
                renderItem={()=>(
                    <View>

                    </View>
                )}
            />
        </Container>
        </>
    );
};

const styles=StyleSheet.create({
    cardImage:{
        width:"100%",
        height:300,
    },
    image:{
        width:350,
        height:200,
        margin:2,
        borderRadius:5
    },
    text:{
        textAlign:"justify"
    }

})

export default DetailDestination;