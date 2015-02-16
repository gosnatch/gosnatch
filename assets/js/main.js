
// url = 'ws://localhost:1313/ws';
//       socket = new WebSocket(url);

//       socket.onmessage = function(msg){
//         console.log(msg)
//         x = JSON.parse(msg.data)
//         ohSnap(x.Message, x.Type)
//       }

//       socket.onopen = function(){
//             console.log('Socket is now open.');
//         };
//         socket.onerror = function (error) {
//             console.error('There was an un-identified Web Socket error');
//         };
//         socket.onclose = function() {
//             console.info( 'Socket is now closed.' );
//         }

//       $(window).on('beforeunload', function(){
//     socket.close();
// });

      // c.onopen = function(){
      //   setInterval(
      //     function(){ send("ping") }
      //   , 1000 )
      // }

$(function () {
  $('[data-toggle="tooltip"]').tooltip({container: "body"})
})