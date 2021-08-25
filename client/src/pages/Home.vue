<template>  
  <div>
    <header>
      <!--Chuong trinh khuyen mai noi bat-->
      <div class="header">Super sale up to 50% off</div>
      <!--Menu-->
      <nav class="container row sticky">
        <div>
            <a href="#"><img src="img/imghome.png" alt="Logo"></a>
        </div>
        <div>
          <ul>
            <li class="dropdown" >
              <a href="#" class="dropbtn">Dropdown1 <i class="fa fa-angle-down"></i></a>
              <div class="dropdown-content">
                  <li><a href="#">Product 1</a></li>
                  <li><a href="#">Product 2</a></li>
              </div>
            </li>
            <li class="dropdown">
              <a href="#" class="dropbtn">Dropdown2 <i class="fa fa-angle-down"></i></a>
              <div class="dropdown-content">
                  <li><a href="#">Product 1</a></li>
                  <li><a href="#">Product 2</a></li>
              </div>
            </li>
            <li><a href="#">Product 1</a></li>
            <li><a href="#">Product 2</a></li>   
            <li> <!--Tim kiem-->
              <div>
                <input type="text" placeholder="Enter keyword to find">               
                <a href="#"><i class="fa fa-search icon"></i></a>
              </div>
            </li>
            <li>
              <div>
                <div class="div-cart">
                    <!-- Docked at the top right corner -->
                    <div class="container__docker">
                        <p class="corner">0</p>
                    </div>
                  <a href="#" class="icon"><i class="fa fa-shopping-bag"></i></a>                 
                </div> 
              </div> 
            </li>             
          </ul>
        </div>
    </nav>

    </header>
    <h1>home</h1>
    <ul>
      <li v-for="p in products" v-bind:key="p.id" float >
        <img src="" width="200px" height="200px">
        <span>{{p.name}}</span>
        <span>{{p.price}}</span>
        <span>{{p.desc}}</span>
        <router-link :to="{name: 'detail', params: {id: p.id }}">Detail</router-link>
        <i class="fa fa-shopping-bag"></i>
      </li>
    </ul>
    <footer>Footer</footer>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Home",
  data() {
    return {
      products: [],
    };
  },
  mounted() {
    this.$nextTick(async function () {
      const { data } = await axios.get("http://localhost:3000/api/products");
      console.log(data)
      this.products =  data
    });
  },
 
};
</script>
<style scoped>
*{
  font-size: 17px;;
}
a{
  text-decoration: none;
}
.header {
  background: #262b2f;
  padding: 0px;
  text-align: center;
  letter-spacing: 0;
  color: #fff;
  padding: 5px;
  font-family: 'Courier New', Courier, monospace;
}
.container {
    display: flex;
    margin-top: 10px;
    margin-left: 70px;
    margin-right: 70px;
   }
.row::after {
    content: "";
    clear: both;
    display: table;
  }
.sticky {
    position: -webkit-sticky;
    position: sticky;
    top: 0;
  }
ul{
   list-style-type: none;   
}
li {
    float: left;
    border: 1px solid;
    padding: 14px 16px;
  }
li a .dropbtn {
    display: block;
    color: rgb(255, 255, 255);
    text-align: center;
    
  }
  
  li a:hover, .dropdown:hover .dropbtn {
    background-color: rgb(255, 255, 255);
  }
  
  li.dropdown {
    display: inline-block;
  }
  
  .dropdown-content {
    display: none;
    position: absolute;
    background-color: #f9f9f9;
    min-width: 160px;
    box-shadow: 0px 8px 16px 0px rgba(0,0,0,0.2);
    z-index: 1;
  }
  
  .dropdown-content a {
    color: rgb(0, 0, 0);
    padding: 12px 16px;
    text-decoration: none;
    display: block;
    text-align: left;
  }
  
  .dropdown-content a:hover {background-color: #8a8a8a;}
  
  .dropdown:hover .dropdown-content {
    display: block;
  }
  /* Style for shopping-cart*/
.div-cart {
  position: relative;
  width: 20px;
}

.container__docker {
    position: absolute;
    transform: translate(50%, -50%);
    padding: 0px;    
}

.corner {
  background:#800000;
  color: white;
  width:17px;
  height:17px;
  border-radius: 20px;
  font-size: 17px;
  /* Center the content */
  align-items: center;
  display: flex;
  justify-content: center;
}
input{
  font-size:17px;
  margin: 0;
}
</style>