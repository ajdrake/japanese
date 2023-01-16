#!/usr/bin/perl -l
use strict;
use warnings;

use LWP::Simple;
use File::Fetch;

# my $url = 'http://www.example.com/file.txt';

sub nihongo() {
    for ( my $i = 0 ; $i <= 23 ; $i++ ) {

        # retrieve the lesson xlsx

        my $lesson =
'https://github.com/SethClydesdale/genki-study-resources/raw/master/resources/tools/wordlist_E-J/lessons-3rd/lesson-'
          . $i . '.xlsx';

        # print $lesson;
        my $lesson_file = 'lesson' . $i . '.xlsx';

        # getstore($lesson, $file);
        my $ff   = File::Fetch->new( uri => $lesson_file );
        my $file = $ff->fetch() or die $ff->error;
    }
}
use Text::Iconv;

# my $converter = Text::Iconv->new("utf-8", "windows-1251");

# Text::Iconv is not really required.
# This can be any object with the convert method. Or nothing.

sub print_nihongo() {

    use Spreadsheet::XLSX;

    for ( my $i = 0 ; $i <= 23 ; $i++ ) {
        my $excel = Spreadsheet::XLSX->new('lessons/lesson-'.$i.'.xlsx');

        foreach my $sheet ( @{ $excel->{Worksheet} } ) {
            $sheet->{MaxRow} ||= $sheet->{MinRow};
            foreach my $row ( $sheet->{MinRow} .. $sheet->{MaxRow} ) {
                $sheet->{MaxCol} ||= $sheet->{MinCol};
                foreach my $col ( $sheet->{MinCol} .. $sheet->{MaxCol} ) {
                    my $cell = $sheet->{Cells}[$row][$col];
                    if ($cell) {
                        printf( "%s\n", $cell->{Val} );
                        
                    }
                }
                printf("\n")
            }
        }
    }
}

print_nihongo();