'use strict';

var gulp = require('gulp');
var browserify = require('browserify');
var babelify = require('babelify');
var source = require('vinyl-source-stream');

// load plugins
//  shorthand method $.pluginName instead of gulpPluginName
var $ = require('gulp-load-plugins')();

gulp.task('styles', function() {
    return gulp.src('src/sass/app.scss')
        .pipe($.sass({
            errLogToConsole: false
        }))
        .pipe($.autoprefixer('last 1 version'))
        .pipe(gulp.dest('assets/css'))
        .pipe($.size())
        .pipe($.notify("CSS-Compilation complete."));
});

gulp.task('browserify', function() {
    return browserify({
        entries: 'src/javascript/page.js',
        debug: true
    })
        .transform(babelify)
        .bundle()
        .pipe(source('bundle.js'))
        .pipe(gulp.dest('assets/js/'));
});

gulp.task('scripts', ['browserify'], function() {
    return gulp.src([
    // other scripts
    './src/javascript/materialize.js',
    './app/dist/bundle.js', ])
        .pipe($.concat('scripts.js'))
    //.pipe($.uglify())
    .pipe(gulp.dest('./app/dist/'))
        .pipe($.size())
        .pipe($.notify("JS-Compilation complete."));
});

gulp.task('images', function() {
    return gulp.src('app/images/**/*')
        .pipe($.cache($.imagemin({
        optimizationLevel: 3,
        progressive: true,
        interlaced: true
    })))
        .pipe(gulp.dest('dist/images'))
        .pipe($.size());
});

gulp.task('clean', function() {
    return gulp.src(['app/styles/main.css', 'dist'], {
        read: false
    }).pipe($.clean());
});

gulp.task('build', ['styles', 'scripts', 'images']);

gulp.task('default', ['clean'], function() {
    gulp.start('build');
});

gulp.task('watch', function() {
    gulp.watch('app/styles/**/*.scss', ['styles']);
    gulp.watch('app/javascript/**/*.js', ['scripts']);
    gulp.watch('app/images/**/*', ['images']);
});