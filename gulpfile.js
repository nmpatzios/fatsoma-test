var gulp = require('gulp'),
child = require('child_process');


gulp.task('build', function () {
child.spawnSync('gb', ['build', 'all'], { stdio: 'inherit' });
});


var server;
gulp.task('run', function () {
if (server) {
  server.kill();
}

server = child.spawn('../../bin/server', { cwd: 'src/server', stdio: 'inherit'});

});

gulp.task('default', ['build', 'run'], function() {
gulp.watch(['src/**/*.go'], ['build', 'run']); // rebuild source when code changes
});
