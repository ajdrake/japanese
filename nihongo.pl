#!/usr/bin/perl -l
use strict;
use warnings;

use LWP::Simple;
use File::Fetch;
# my $url = 'http://www.example.com/file.txt';
sub nihongo() 
{
    for ( my $i = 0; $i <= 23; $i++ ) {
        # retrieve the lesson xlsx
        
        my $lesson = 'https://github.com/SethClydesdale/genki-study-resources/raw/master/resources/tools/wordlist_E-J/lessons-3rd/lesson-' . $i . '.xlsx';
        # print $lesson;
        my $lesson_file = 'lesson'.$i.'.xlsx';
        # getstore($lesson, $file);
        my $ff = File::Fetch->new(uri => $lesson_file);
        my $file = $ff->fetch() or die $ff->error;
    }
}

nihongo();