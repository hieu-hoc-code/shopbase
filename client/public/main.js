var app = new Vue({
  el: '#app',
  data: {
    message: 'Nội dung quảng cáo'
  }
})

Vue.component('todo-item', {
  props: ['todo'],
  template: '<a href={{todo.href}}>{{ todo.text }}</a>'
})
Vue.component('todo-mnitem', {
  props: ['todo'],
  template: '<li><a href={{todo.href}}>{{ todo.text }}</a></li>'
})

var app1 = new Vue({
  el:'#menu',
  data:{
    List1: [
      { href:'#', text: 'Link 1' },
      { href:'#', text: 'Link 1' },
      { href:'#', text: 'Link 1' }
    ],
    List2: [
      { href:'#', text: 'Link 2' },
      { href:'#', text: 'Link 2' },
      { href:'#', text: 'Link 2' }
    ],
  }
})
var app2 = new Vue({
  el:"Login",
  
})
