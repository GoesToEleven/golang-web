var gulp = require('gulp');
var sass = require('gulp-sass');
var jshint = require('gulp-jshint');
var webserver = require('gulp-webserver');

gulp.task('webserver', function () {
    gulp.src('builds/development')
        .pipe(webserver({
            livereload: true,
            directoryListing: false,
            open: true
        }));
});

gulp.task('sass', function () {
    gulp.src('./components/sass/style.sass')
        .pipe(sass().on('error', sass.logError))
        .pipe(gulp.dest('builds/development/css'));
});

gulp.task('jshint', function() {
    return gulp.src('./builds/development/**/*.js')
        .pipe(jshint())
        .pipe(jshint.reporter('default'));
});

gulp.task('watch', function () {
    gulp.watch('./components/sass/style.sass', ['sass']);
});

gulp.task('default', ['webserver', 'sass', 'jshint', 'watch']);