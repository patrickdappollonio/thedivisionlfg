var gulp = require('gulp');
var concat = require('gulp-concat');
var uglify = require('gulp-uglify');
var cleanCSS = require('gulp-clean-css');
var autoprefixer = require('gulp-autoprefixer');
var stripCssComments = require('gulp-strip-css-comments');
var imagemin = require('gulp-imagemin');
var stylus = require('gulp-stylus');
var es = require('event-stream');
var nib = require('nib');
var del = require('del');

gulp.task('cleanjs', function() {
	return del(['assets/build/js'])
});
gulp.task('cleancss', function() {
	return del(['assets/build/css'])
});
gulp.task('cleanimg', function() {
	return del(['assets/build/img'])
});

gulp.task('js', ['cleanjs'], function() {
	return gulp.src([
			'assets/js/jquery.min.js',
			'assets/js/bootstrap.min.js',
			'assets/js/bootbox.js',
			'assets/js/javascripts.js',
		])
		.pipe(uglify())
		.pipe(concat('build.min.js'))
		.pipe(gulp.dest('assets/build/js'));
});

gulp.task('css', ['cleancss'], function() {
	return gulp.src('assets/css/*.styl')
		.pipe(stylus({
			'include css': true,
			use: nib(),
			import: ['nib'],
			compress: true
		}))
		.pipe(cleanCSS())
		.pipe(autoprefixer('last 2 version', 'safari 5', 'ie 8', 'ie 9'))
		.pipe(stripCssComments())
		.pipe(concat('build.min.css'))
		.pipe(gulp.dest('assets/build/css'));
});

gulp.task('img', ['cleanimg'], function() {
	return gulp.src(['assets/img/*.jpg', 'assets/img/*.gif', 'assets/img/*.png'])
		.pipe(imagemin({
			optimizationLevel: 5
		}))
		.pipe(gulp.dest('assets/build/img'));
});

gulp.task('watch', function() {
	gulp.watch('assets/js/*.*', ['js']);
	gulp.watch('assets/css/*.*', ['css']);
	gulp.watch(['assets/img/*.jpg', 'assets/img/*.gif', 'assets/img/*.png'], ['img']);
});

gulp.task('default', ['watch', 'js', 'css', 'img']);